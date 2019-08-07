package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type contextKey string

const UserContextKey contextKey = "user"

type User struct {
	FirebaseUID string `db:"firebase_uid"`
	DisplayName string `db:"display_name"`
	Email       string `db:"email"`
	PhotoURL    string `db:"photo_url"`
}

func SetUserToContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, UserContextKey, u)
}

func GetUserFromContext(ctx context.Context) (*User, error) {
	v := ctx.Value(UserContextKey)
	user, ok := v.(*User)
	if !ok {
		return nil, errors.New("user not found from context value")
	}
	return user, nil
}

func GetUser(dbx *sqlx.DB, uid string) (*User, error) {
	var u User
	if err := dbx.Get(&u, `
select firebase_uid, display_name, email, photo_url from user where firebase_uid = ? limit 1
	`, uid); err != nil {
		return nil, err
	}
	return &u, nil
}

func SyncUser(db *sqlx.DB, u *User) (sql.Result, error) {
	return db.Exec(`
INSERT INTO user (firebase_uid, display_name, email, photo_url)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY
UPDATE display_name = ?, email = ?, photo_url = ?, utime = NOW()
`, u.FirebaseUID, u.DisplayName, u.Email, u.PhotoURL, u.DisplayName, u.Email, u.PhotoURL)
}
