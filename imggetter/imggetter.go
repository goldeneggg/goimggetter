package imggetter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

const (
	TIMEOUT_SEC = 60
	IMG_DIR     = "img/"
	USER_AGENT  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
)

type imgGetter struct {
	query   string
	page    int
	referer string
}

type scraper interface {
	search() ([]string, error)
	getSrcs(searchResult string) ([]string, error)
	getNormalizedSrcs(srcs []string) []string
	getReferer() string
}

var scr scraper

func Download(site string, query string, page int, concurrency int, debug bool) error {
	// prepare logger
	prepareLogger(debug)

	// configure scraper
	configureScraper(site, query, page)
	if scr == nil {
		return fmt.Errorf("Invalid site: %s\n", site)
	}

	// search
	l.debug("Start query:", query, "page:", page)
	results, err := scr.search()
	if err != nil {
		return err
	}
	l.debug("Results:", results, "count:", len(results))

	// get img srcs and download
	for _, result := range results {
		// mkdir for download history
		if mkHistDir(result) {
			// get image srcs
			srcs, err := scr.getSrcs(result)
			if err != nil {
				return err
			}
			l.debug("Original srcs:", srcs)

			// normalize srcs
			nSrcs := scr.getNormalizedSrcs(srcs)
			l.debug("Normalized srcs:", nSrcs, "in page", result)

			// save by concurrency proccessing
			from := 0
			for from < len(nSrcs) {
				var to int
				if len(nSrcs[from:]) > concurrency {
					to = from + concurrency
				} else {
					to = from + len(nSrcs[from:])
				}

				var wg sync.WaitGroup
				for i, ns := range nSrcs[from:to] {
					wg.Add(1)
					go func(target string, idx int) {
						defer wg.Done()
						save(target)
						l.debug("Done:", idx, target)
					}(ns, i)
				}
				wg.Wait()

				from += concurrency
			}
		}
	}

	return nil
}

func configureScraper(site string, query string, page int) {
	if site == "flickr" {
		scr = newFlickrImgGetter(query, page)
	}
}

func httpGet(url string) (*http.Response, error) {
	// client object
	client := &http.Client{Timeout: time.Duration(TIMEOUT_SEC * time.Second)}

	// request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// user-agent and referer
	req.Header.Add("User-Agent", USER_AGENT)
	if referer := scr.getReferer(); len(referer) > 0 {
		req.Header.Add("Referer", referer)
	}

	// http access
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Invalid StatusCode: %v, Status: %v", resp.StatusCode, resp.Status)
		resp.Body.Close()
		return nil, err
	}

	return resp, nil
}

func mkHistDir(result string) bool {
	os.Mkdir(IMG_DIR, 0755)

	dir := IMG_DIR + path.Base(path.Dir(result)) + "_" + path.Base(result)
	if err := os.Mkdir(dir, 0755); err != nil {
		l.err(err)
		return false
	}

	return true
}

func save(src string) bool {
	// http access
	resp, err := httpGet(src)
	if err != nil {
		l.err(err)
		return false
	}
	defer resp.Body.Close()

	// file open
	fileName := path.Base(src)
	f, err := os.OpenFile(IMG_DIR+fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		l.err(err)
		return false
	}
	defer f.Close()

	// save
	written, err := io.Copy(f, resp.Body)
	if err == nil {
		l.debug("Saved:", fileName, "(", written, " bytes)")
		return true
	} else {
		l.err(err)
		return false
	}
}
