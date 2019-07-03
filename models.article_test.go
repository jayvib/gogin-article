package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	articles := getAllArticles()
	if len(articles) != len(articleList) {
		t.Error("number of items doesn't match")
	}
}
