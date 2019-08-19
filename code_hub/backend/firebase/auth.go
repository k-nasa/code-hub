package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func InitAuthClient() (*auth.Client, error) {
	fb, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return fb.Auth(context.Background())
}
