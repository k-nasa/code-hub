package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"net/http"
	"strconv"
)

type Article struct {
	dbx *sqlx.DB
}

func NewArticle(dbx *sqlx.DB) *Article {
	return &Article{dbx: dbx}
}

func (a *Article) Root(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	articles, err := model.ListArticle(a.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, articles, nil
}

func (a *Article) Get(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	article, err := model.GetArticle(a.dbx, aid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusCreated, article, nil
}

func (a *Article) New(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	article := &model.Article{}
	err := decoder.Decode(&article)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	result, err := model.Insert(a.dbx, article)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	article.ID = id
	return http.StatusCreated, article, nil
}

func (a *Article) Edit(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	article := &model.Article{}
	err := decoder.Decode(&article)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusCreated, nil, nil
}
