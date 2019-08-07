package repository

import (
	"database/sql"

	model2 "github.com/voyagegroup/treasure-app/model"

	"github.com/jmoiron/sqlx"
)

func GetUser(db *sqlx.DB, uid string) (*model2.User, error) {
	var u model2.User
	if err := db.Get(&u, `
select firebase_uid, display_name, email, photo_url from user where firebase_uid = ? limit 1
	`, uid); err != nil {
		return nil, err
	}
	return &u, nil
}

func SyncUser(db *sqlx.DB, u *model2.User) (sql.Result, error) {
	return db.Exec(`
INSERT INTO user (firebase_uid, display_name, email, photo_url)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY
UPDATE display_name = ?, email = ?, photo_url = ?, utime = NOW()
`, u.FirebaseUID, u.DisplayName, u.Email, u.PhotoURL, u.DisplayName, u.Email, u.PhotoURL)
}
