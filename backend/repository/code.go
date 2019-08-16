package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/model"
)

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