package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/httputil"
	"github.com/k-nasa/code-hub/model"
	"github.com/k-nasa/code-hub/service"
)

type Code struct {
	dbx *sqlx.DB
}

func NewCode(dbx *sqlx.DB) *Code {
	return &Code{dbx: dbx}
}

func (c *Code) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newCode := &model.Code{}

	if err := json.NewDecoder(r.Body).Decode(&newCode); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newCode.UserID = &user.ID

	codeService := service.NewCodeService(c.dbx)

	newCode, err = codeService.Create(newCode)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, newCode, nil
}
