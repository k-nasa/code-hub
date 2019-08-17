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
	"github.com/k-nasa/code-hub/service"
)

type Code struct {
	db *sqlx.DB
}

func NewCode(db *sqlx.DB) *Code {
	return &Code{db: db}
}

func (c *Code) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	codes, err := repository.AllCodes(c.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, codes, nil
}

func (c *Code) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	code, err := repository.FindCode(c.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, code, nil
}

func (c *Code) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newCode := &model.Code{}

	if err := json.NewDecoder(r.Body).Decode(&newCode); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if newCode.Title == "" || newCode.Body == "" {
		return http.StatusBadRequest, nil, nil
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newCode.UserID = &user.ID

	codeService := service.NewCodeService(c.db)

	newCode, err = codeService.Create(newCode)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, newCode, nil
}

func (c *Code) IndexWithUser(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	code, err := repository.AllCodesWithUser(c.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, code, nil
}

func (c *Code) ShowUserCode(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	codeService := service.NewCodeService(c.db)
	code, err := codeService.FindUserCode(aid)

	return http.StatusOK, code, nil
}
