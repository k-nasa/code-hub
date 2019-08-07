package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func GetUser(db *sqlx.DB, uid string) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `
select firebase_uid, display_name, email, photo_url from user where firebase_uid = ? limit 1
	`, uid); err != nil {
		return nil, err
	}
	return &u, nil
}

func SyncUser(db *sqlx.DB, u *model.User) (sql.Result, error) {
	return db.Exec(`
INSERT INTO user (firebase_uid, display_name, email, photo_url)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY
UPDATE display_name = ?, email = ?, photo_url = ?, utime = NOW()
`, u.FirebaseUID, u.DisplayName, u.Email, u.PhotoURL, u.DisplayName, u.Email, u.PhotoURL)
}
