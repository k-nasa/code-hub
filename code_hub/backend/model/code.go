package model

import "time"

type Code struct {
	ID     int64  `db:"id" json:"id"`
	UserID *int64 `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
	// TODO あとからenumっぽいのにしたい
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CodeWithUser struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Body        string `db:"body" json:"body"`
	FirebaseUID string `db:"firebase_uid" json:"firebase_uid"`
	UserID      *int64 `db:"user_id" json:"user_id"`
	IconUrl     string `db:"icon_url" json:"icon_url"`
	Username    string `db:"username" json:"username"`
	// TODO あとからenumっぽいのにしたい
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
