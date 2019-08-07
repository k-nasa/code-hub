package model

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Article struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
}

func ListArticle(dbx *sqlx.DB) ([]Article, error) {
	var a []Article
	if err := dbx.Select(&a, `SELECT id, title, body FROM article`); err != nil {
		return nil, err
	}
	return a, nil
}

func GetArticle(dbx *sqlx.DB, id int64) (*Article, error) {
	a := Article{}
	if err := dbx.Get(&a, `
SELECT id, title, body FROM article WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func Insert(db *sqlx.DB, a *Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (title, body) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(a.Title, a.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Update(db *sqlx.DB, id int64, a *Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article SET title = ?, body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, id)
}
