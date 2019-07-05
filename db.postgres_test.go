package main

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestSetupGormDB(t *testing.T) {
	db, err := setupGormDB(true)
	if err != nil {
		t.Fatal(err)
	}

	inst := 5
	for i := 0; i < inst; i++ {
		__db, err := setupGormDB(true)
		if err != nil {
			t.Fail()
		}
		if !reflect.DeepEqual(db, __db) {
			t.Error("first instance is not the same pointer with the succedding instance")
		}
	}
}

func TestSetupTable(t *testing.T) {
	db, err := setupGormDB(true)
	if err != nil {
		t.Fatal(err)
	}
	setupTable(db)
}

func TestDBPutArticle(t *testing.T) {
	db, err := setupGormDB(true)
	if err != nil {
		t.Fatal(err)
	}
	initMigrationWithDefaultDB()
	pdb := &postgresDB{db: db}
	article1 := &Article{
		Title:   "test1",
		Content: "this is a testing 2",
	}
	pdb.putArticle(article1)

	t.Log(article1.ID)
	t.Log(article1.CreatedAt)

	var articles []*Article
	articles, _ = pdb.getAllArticles()
	if len(articles) != 1 {
		t.Error("after putting 1 item successfully, there should be 1 article in the table")
	}
}

func TestDBGetArticle(t *testing.T) {
	db, err := setupGormDB(true)
	if err != nil {
		t.Fatal(err)
	}
	initMigrationWithDefaultDB()
	article1 := &Article{
		Title:   "test1",
		Content: "this is a testing 2",
	}
	article2 := &Article{
		Title:   "test2",
		Content: "this is a testing 2",
	}
	pdb := &postgresDB{db: db}
	pdb.putArticle(article1)
	pdb.putArticle(article2)

	article1query := &Article{Model: gorm.Model{ID: 1}}
	article1query, err = pdb.getArticle(article1query)
	if err != nil {
		t.Fatal(err)
	}
	if article1.Title != article1query.Title {
		t.Error("Title from the original is not the same with the title of the article result")
	}

}
