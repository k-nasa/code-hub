package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/model"
)

func CreateComment(db *sqlx.DB, comment *model.Comment) (sql.Result, error) {
	return db.Exec(`
insert into
comments(user_id, code_id, body)
values(?, ?, ?)
`, comment.UserID, comment.CodeID, comment.Body)
}

func AllComments(db *sqlx.DB, codeID int64) ([]model.CommentWithUser, error) {
	c := make([]model.CommentWithUser, 0)
	if err := db.Select(&c, `select comments.*, username, icon_url from comments inner join users on user_id = users.id where code_id = ?`, codeID); err != nil {
		return nil, err
	}
	return c, nil
}

func FindComment(db *sqlx.DB, id int64) (*model.CommentWithUser, error) {
	code := model.CommentWithUser{}

	if err := db.Get(&code, `SELECT codes.*, icon_url, username FROM codes inner join users on codes.user_id = users.id WHERE id = ? LIMIT 1`, id); err != nil {
		return nil, err
	}
	return &code, nil
}
