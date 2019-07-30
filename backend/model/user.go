package model

import (
	"net/http"
)

const UserContextKey = "user"

type User struct {
	FirebaseUID string
	DisplayName string
	Email       string
	PhotoURL    string
}

func GetUserFromContext(req *http.Request) User {
	return req.Context().Value(UserContextKey).(User)
}
