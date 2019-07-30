package model

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Comment struct {
	ID      int64      `db:"id"`
	Uid     string     `db:"user_uid"`
	Content string     `db:"content"`
	Ctime   *time.Time `db:"ctime"`
}

func GetComments(dbx *sqlx.DB, uid string) ([]Comment, error) {
	comments := make([]Comment, 0)
	if err := dbx.Select(&comments, `
select * from comment where user_uid = ?
	`, uid); err != nil {
		return nil, err
	}
	return comments, nil
}
