package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/model"
)

func AllCodes(db *sqlx.DB) ([]model.Code, error) {
	c := make([]model.Code, 0)
	if err := db.Select(&c, `SELECT * FROM codes`); err != nil {
		return nil, err
	}
	return c, nil
}

func AllCodesWithUser(db *sqlx.DB) ([]model.CodeWithUser, error) {
	c := make([]model.CodeWithUser, 0)
	if err := db.Select(&c, `select codes.id, title, body, status, codes.created_at, codes.updated_at, user_id, username from codes inner join users on users.id = codes.user_id order by codes.updated_at desc`); err != nil {
		return nil, err
	}
	return c, nil
}

func FindCode(db *sqlx.DB, id int64) (*model.Code, error) {
	code := model.Code{}

	if err := db.Get(&code, `SELECT * FROM codes WHERE id = ? LIMIT 1`, id); err != nil {
		return nil, err
	}
	return &code, nil
}

func CreateCode(db *sqlx.Tx, code *model.Code) (sql.Result, error) {
	return db.Exec(`insert into codes(user_id, title, body, status) values(?, ?, ?, ?)`,
		code.UserID, code.Title, code.Body, code.Status)
}
