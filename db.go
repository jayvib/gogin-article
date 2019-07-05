package main

import "errors"

type DB interface {
	getAllArticles() ([]Article, error)
	getArticle(id int) (Article, error)
	putArticle(*Article) error
}

// A Data store using protocol buffer
type protoDB struct {
}

func (pdb *protoDB) getAllArticles() (*article, error) {
	return nil, errors.New("not implemented")
}

func (pdb *protoDB) getArticle(id string) (*article, error) {
	return nil, errors.New("not implemented")
}
