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
