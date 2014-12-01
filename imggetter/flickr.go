package imggetter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type flickrImgGetter imgGetter

func newFlickrImgGetter(q string, p int) *flickrImgGetter {
	return &flickrImgGetter{query: q, page: p, referer: ""}
}

func (fg *flickrImgGetter) search() ([]string, error) {
	results := make([]string, 0, 60)

	searchUrl := fmt.Sprintf("https://www.flickr.com/search/?text=%s&content_type=1&media=photos&adv=1&sort=date-posted-desc&page=%d", fg.query, fg.page)
	doc, err := goquery.NewDocument(searchUrl)
	if err != nil {
		l.err("Error NewDocument", err)
		return results, err
	}

	doc.Find(".photo-click").Each(func(i int, s *goquery.Selection) {
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

	doc.Find("#image-src").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("href"); ok {
			srcs = append(srcs, src)
		}
	})

	return srcs, nil
}

func (fg *flickrImgGetter) getNormalizedSrcs(srcs []string) []string {
	return []string{strings.Replace(srcs[0], "_m", "_b", 1)}
}

func (fg *flickrImgGetter) getReferer() string {
	return fg.referer
}
