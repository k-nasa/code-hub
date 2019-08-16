package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/model"
)

func GetUser(db *sqlx.DB, uid string) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `
SELECT id, firebase_uid, username, email, icon_url FROM users WHERE firebase_uid = ? LIMIT 1
	`, uid); err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserById(db *sqlx.DB, id int64) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `
SELECT id, firebase_uid, username, email, icon_url FROM users WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &u, nil
}

func SyncUser(db *sqlx.DB, fu *model.FirebaseUser) (sql.Result, error) {
	return db.Exec(`
INSERT INTO users (firebase_uid, username, email, icon_url)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY
UPDATE username = ?, email = ?, icon_url = ?
`, fu.FirebaseUID, fu.DisplayName, fu.Email, fu.PhotoURL, fu.DisplayName, fu.Email, fu.PhotoURL)
}
