package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/repository"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	users, err := repository.AllCodes(c.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}
