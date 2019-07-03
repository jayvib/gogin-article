package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplate bool) *gin.Engine {
	router := gin.Default()
	if withTemplate {
		router.LoadHTMLGlob("templates/*")
	}
	return router
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if !f(rec) {
		t.Fail()
	}
}

func saveList() {
	tmpArticleList = articleList
}

func restoreList() {
	articleList = tmpArticleList
}
