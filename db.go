package main

import "errors"

type DB interface {
	getAllArticles() (*[]article, error)
	getArticle(id string) (*article, error)
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
