package main

import (
	"fmt"
	"sync"

	"github.com/jayvib/gogin-article/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	USERNAME = "postgres"
	PASSWORD = "qwerty"
	DBNAME   = "articledb"
)

// Use singleton to avoid establishing new connection
var (
	once sync.Once // will initialize DB only once
	_db  *gorm.DB
)

type postgresDB struct {
	db *gorm.DB
}

func (pdb *postgresDB) getAllArticles() ([]*Article, error) {
	var articles []*Article
	if err := pdb.db.Find(&articles).Error; err != nil {
		return nil, errors.Wrap(err, "error while getting all the articles")
	}
	return articles, nil
}

func (pdb *postgresDB) getArticle(a *Article) (*Article, error) {
	if err := pdb.db.First(a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ItemNotFound.Newf("article %v not found", a)
		}
	}
	return a, nil
}

func (pdb *postgresDB) putArticle(a *Article) error {
	if err := pdb.db.Create(a).Error; err != nil {
		return errors.Wrap(err, "error while putting the article")
	}
	return nil
}

func setupGormDB(isDebug bool) (*gorm.DB, error) {
	datasource := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		USERNAME, DBNAME, PASSWORD,
	)
	var err error
	once.Do(func() {
		_db, err = gorm.Open("postgres", datasource)
		if isDebug {
			_db = _db.Debug()
		}
	})
	if err != nil {
		return nil, errors.Wrap(err, "something wrong while opening the database")
	}
	return _db, nil
}

func setupTable(db *gorm.DB) {
	db.DropTableIfExists(&Article{})
	db.AutoMigrate(&Article{})
}

func initMigrationWithDefaultDB() error {
	db, err := setupGormDB(true)
	if err != nil {
		return err
	}
	setupTable(db)
	return nil
}
