package model

import "time"

type User struct {
	ID          int64     `db:"id" json:"id"`
	FirebaseUID string    `db:"firebase_uid" json:"firebase_uid"`
	Email       string    `db:"email" json:"email"`
	Username    string    `db:"username" json:"username"`
	IconUrl     string    `db:"icon_url" json:"icon_url"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type UserCodes struct {
	User  User    `json:"user"`
	Codes []*Code `json:"codes"`
}
