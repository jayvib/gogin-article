package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	USERNAME = "postgres"
	PASSWORD = "qwerty"
	DBNAME   = "articledb"
)

type postgresDB struct {
	db *gorm.DB
}

func (pdb *postgresDB) getAllArticles() ([]*article, error) {
	return nil, nil
}

func (pdb *postgresDB) getArticle(id int) (*article, error) {
	return nil, nil
}

func (pdb *postgresDB) putArticle(a *article) error {
	pdb.db.Create(a)
	return nil
}

func setupGormDB(isDebug bool) (*gorm.DB, error) {
	datasource := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		USERNAME, DBNAME, PASSWORD,
	)
	db, err := gorm.Open("postgres", datasource)
	if err != nil {
		return nil, fmt.Errorf("error while opening the database, %v", err)
	}

	if isDebug {
		db = db.Debug()
	}
	return db, nil
}

func setupTable(db *gorm.DB) {
	db.DropTableIfExists(&article{})
	db.CreateTable(&article{})
}
