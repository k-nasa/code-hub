package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	model2 "github.com/voyagegroup/treasure-app/model"
)

func AllArticle(db *sqlx.DB) ([]model2.Article, error) {
	var a []model2.Article
	if err := db.Select(&a, `SELECT id, title, body FROM article`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindArticle(db *sqlx.DB, id int64) (*model2.Article, error) {
	a := model2.Article{}
	if err := db.Get(&a, `
SELECT id, title, body FROM article WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateArticle(db *sqlx.Tx, a *model2.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (title, body) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body)
}

func UpdateArticle(db *sqlx.Tx, id int64, a *model2.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article SET title = ?, body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, id)
}

func DestroyArticle(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
