package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/domain/model"
	"github.com/voyagegroup/treasure-app/domain/service"
	"github.com/voyagegroup/treasure-app/httputil"
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
	article := &model.Article{}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
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
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqArticle := &model.Article{}
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleService := service.NewArticleService(a.dbx)
	_, err = articleService.Update(aid, reqArticle)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusNoContent, nil, nil
}
