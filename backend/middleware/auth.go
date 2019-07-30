package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/voyagegroup/treasure-app/model"
	"log"
	"net/http"
)

const (
	bearer = "Bearer"
)

type AuthMiddleware struct {
	client *auth.Client
}

func NewAuthMiddleware(client *auth.Client) *AuthMiddleware {
	return &AuthMiddleware{
		client: client,
	}
}

func (auth *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken, err := getTokenFromHeader(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := auth.client.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			log.Printf(err.Error())
			http.Error(w, "Failed to verify token", http.StatusForbidden)
			return
		}
		user, err := auth.client.GetUser(r.Context(), token.UID)

		if err != nil {
			log.Printf(err.Error())
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), model.UserContextKey, model.User{
			FirebaseUID: user.UID,
			Email:       user.Email,
			DisplayName: user.DisplayName,
			PhotoURL:    user.PhotoURL,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenFromHeader(req *http.Request) (string, error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("authorization header not found")
	}

	l := len(bearer)
	if len(header) > l+1 && header[:l] == bearer {
		return header[l+1:], nil
	}

	return "", fmt.Errorf("authorization header format must be 'Bearer {token}'")
}
