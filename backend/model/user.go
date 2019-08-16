package model

import "time"

type User struct {
	ID          int64     `db:"id"`
	FirebaseUID string    `db:"firebase_uid"`
	Email       string    `db:"email"`
	Username    string    `db:"username"`
	IconUrl     string    `db:"icon_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
