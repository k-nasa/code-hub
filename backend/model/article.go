package model

type Article struct {
	ID    string `db:"id"`
	Title string `db:"id"`
	Body  string `db:"body"`
}
