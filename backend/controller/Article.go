package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Article struct {
	dbx *sqlx.DB
}

func NewArticle(dbx *sqlx.DB) *Article {
	return &Article{dbx: dbx}
}

func (a *Article) Root(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, nil, nil
}
