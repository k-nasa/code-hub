package model

import "time"

type Comment struct {
	ID        int64     `db:"id", json:"id"`
	UserID    *int64    `db:"user_id" json:"user_id"`
	CodeID    int64     `db:"code_id" json:"code_id"`
	Body      string    `db:"body": json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CommentWithUser struct {
	ID       int64  `db:"id" json:"id"`
	UserID   *int64 `db:"user_id" json:"user_id"`
	CodeID   int64  `db:"code_id" json:"code_id"`
	Body     string `db:"body" json:"body"`
	IconUrl  string `db:"icon_url" json:"icon_url"`
	Username string `db:"username" json:"username"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
