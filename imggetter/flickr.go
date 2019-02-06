package imggetter

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type flickrImgGetter imgGetter

func newFlickrImgGetter(q string, p int) *flickrImgGetter {
	return &flickrImgGetter{query: q, page: p, referer: ""}
}

func (fg *flickrImgGetter) search() ([]string, error) {
	results := make([]string, 0, 60)

	//searchUrl := fmt.Sprintf("https://www.flickr.com/search/?text=%s&content_type=1&media=photos&adv=1&sort=date-posted-desc&page=%d", fg.query, fg.page)
	searchUrl := fmt.Sprintf("https://www.flickr.com/search/?text=%s&view_all=1", fg.query)
	l.debug("searchUrl:", searchUrl)

	// http access
	resp, err := httpGet(searchUrl)
	if err != nil {
		l.err(err)
		return results, err
	}
	defer resp.Body.Close()
	l.debug("resp.Body:", resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		l.err("Error NewDocument", err)
		return results, err
	}
	html, err := doc.Html()
	if err != nil {
		l.err("Error doc.Html()", err)
		return results, err
	}
	l.debug("html:", html)

	//doc.Find(".photo-click").Each(func(i int, s *goquery.Selection) {
	doc.Find(".overlay").Each(func(i int, s *goquery.Selection) {
		l.debug("s: ", s)
		if result, ok := s.Attr("href"); ok {
			results = append(results, result)
		}
	})

	return results, nil
}

func (fg *flickrImgGetter) getSrcs(searchResult string) ([]string, error) {
	srcs := make([]string, 0, 1)

	// searchResult is absolutePath
	doc, err := goquery.NewDocument(fmt.Sprintf("https://www.flickr.com%s", searchResult))
	if err != nil {
		l.err("Error NewDocument", err)
		return srcs, err
	}

	//doc.Find("#image-src").Each(func(i int, s *goquery.Selection) {
	doc.Find(".low-res-photo").Each(func(i int, s *goquery.Selection) {
		//if src, ok := s.Attr("href"); ok {
		if src, ok := s.Attr("src"); ok {
			srcs = append(srcs, src)
		}
	})

	return srcs, nil
}

func (fg *flickrImgGetter) getNormalizedSrcs(srcs []string) []string {
	//return []string{strings.Replace(srcs[0], "_m", "_b", 1)}
	return srcs
}

func (fg *flickrImgGetter) getReferer() string {
	return fg.referer
}
