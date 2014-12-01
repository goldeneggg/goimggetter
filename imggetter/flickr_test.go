package imggetter

import (
	"regexp"
	"strings"
	"testing"
)

const (
	SEARCH_RESULT_REGEXP = "1[0-9]{10}/$"
	SRC_SUFFIX           = "_m.jpg"
	NSRC_SUFFIX          = "_b.jpg"
)

var (
	gettersNormal = []*flickrImgGetter{
		&flickrImgGetter{"tokyo", 1, ""},
		&flickrImgGetter{"%E6%9D%B1%E4%BA%AC", 999, ""},
	}

	gettersEmpty = []*flickrImgGetter{
		&flickrImgGetter{"aaaaaaaaaaaaaaaaxxxxxxxxxxxxxxx", 1, ""},
	}
)

func TestFlickrNormal(t *testing.T) {
	for _, sc := range gettersNormal {
		results, err := sc.search()

		if err != nil {
			t.Errorf("search error: %#v, query: %s, page: %d", err, sc.query, sc.page)
		}

		t.Log("results", results)
		resultsLen := len(results)
		if resultsLen == 0 {
			t.Errorf("search result is nothing, query: %s, page: %d", sc.query, sc.page)
		} else {
			testGetSrcs(results, sc, t)
		}
	}
}

func testGetSrcs(results []string, sc scraper, t *testing.T) {
	for _, result := range results {
		t.Log("result", result)
		matched, err := regexp.MatchString(SEARCH_RESULT_REGEXP, result)
		if err != nil {
			t.Errorf("invalid result URL: %#v, matched: %s", result, matched)
		}

		srcs, err := sc.getSrcs(result)
		if err != nil {
			t.Errorf("GetSrcs error: %#v, result: %s", err, result)
		}
		srcsLen := len(srcs)
		if srcsLen != 1 {
			t.Errorf("srcs do not have 1 length, result: %s", result)
		}
		if !strings.HasSuffix(srcs[0], SRC_SUFFIX) {
			t.Errorf("src[%s] do not have expected suffix[%s]", srcs[0], SRC_SUFFIX)
		}

		nSrcs := sc.getNormalizedSrcs(srcs)
		nSrcsLen := len(nSrcs)
		if nSrcsLen != 1 {
			t.Errorf("nSrcs do not have 1 length, nSrcs: %s", nSrcs)
		}
		if !strings.HasSuffix(nSrcs[0], NSRC_SUFFIX) {
			t.Errorf("nsrc[%s] do not have expected suffix[%s]", nSrcs[0], NSRC_SUFFIX)
		}
	}
}

func TestFlickrEmpty(t *testing.T) {
	for _, sc := range gettersEmpty {
		results, _ := sc.search()

		if len(results) > 0 {
			t.Errorf("returned non-zero search results: %#v, query: %s, page: %d", results, sc.query, sc.page)
		}
	}
}

func BenchmarkFlickrNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc := gettersNormal[0]
		results, _ := sc.search()
		srcs, _ := sc.getSrcs(results[0])
		sc.getNormalizedSrcs(srcs)
	}
}
