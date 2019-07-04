package main

import "testing"

func TestSetupGormDB(t *testing.T) {
	_, err := setupGormDB(true)
	if err != nil {
		t.Fatal(err)
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
	pdb := &postgresDB{db: db}
	article1 := &article{
		Title:   "test1",
		Content: "this is a testing",
	}
	pdb.putArticle(article1)
}
