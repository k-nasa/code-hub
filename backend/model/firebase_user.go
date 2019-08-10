package model

type FirebaseUser struct {
	FirebaseUID string `db:"firebase_uid"`
	DisplayName string `db:"display_name"`
	Email       string `db:"email"`
	PhotoURL    string `db:"photo_url"`
}
