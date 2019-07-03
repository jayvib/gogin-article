package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		if w.Code != http.StatusOK {
			return false
		}

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}

		if strings.Index(string(p), "<title>Home Page</title>") < 0 {
			return false
		}
		return true
	})

}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest(http.MethodGet, "/article/view/1", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		if w.Code != http.StatusOK {
			return false
		}
		content, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		if strings.Index(string(content), "Article 1") < 0 {
			return false
		}
		return true

	})
}
