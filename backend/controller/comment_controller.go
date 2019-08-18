package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/httputil"
	"github.com/k-nasa/code-hub/model"
	"github.com/k-nasa/code-hub/repository"
)

type Comment struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *Comment {
	return &Comment{db: db}
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newComment := &model.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newComment.UserID = &user.ID

	result, err := repository.CreateComment(c.db, newComment)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// FIXME ここらへんはserviceに切り出す!!
	id, err := result.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	comment, err := repository.FindComment(c.db, id)

	return http.StatusCreated, comment, nil
}

func (c *Comment) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	codeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	comments, err := repository.AllComments(c.db, codeID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, comments, nil
}
