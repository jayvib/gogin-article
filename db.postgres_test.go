package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/jayvib/gogin-article/errors"
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
	err = initMigrationWithDefaultDB()
	if err != nil {
		t.Fatal(err)
	}
	pdb := &postgresDB{db: db}
	article1 := &Article{
		Title:   "test1",
		Content: "this is a testing 2",
	}
	err = pdb.putArticle(article1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(article1.ID)
	t.Log(article1.CreatedAt)

	var articles []*Article
	articles, err = pdb.getAllArticles()
	if err != nil {
		log.Fatal(err)
	}
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

	notexistArticle := &Article{Model: gorm.Model{ID: 1000}}
	_, err = pdb.getArticle(notexistArticle)
	if err == nil {
		t.Error("expecting an error but nothing return")
	}

	if !errors.IsCustomError(err) {
		t.Error("expecting an customError type but it doesn't return the correct on")
	} else {
		if et := errors.GetErrorType(err); et != errors.ItemNotFound {
			t.Errorf("expecting an error type 'ItemNotFound' but got %v", et)
		}
	}
}

func TestInitMigration(t *testing.T) {
	err := initMigrationWithDefaultDB()
	if err != nil {
		t.Fatal(err)
	}

}
