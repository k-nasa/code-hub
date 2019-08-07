package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/voyagegroup/treasure-app/domain/model"
	"github.com/voyagegroup/treasure-app/domain/repository"
	"github.com/voyagegroup/treasure-app/httputil"

	"firebase.google.com/go/auth"
	"github.com/jmoiron/sqlx"
)

const (
	bearer = "Bearer"
)

type AuthMiddleware struct {
	client *auth.Client
	db     *sqlx.DB
}

func NewAuthMiddleware(client *auth.Client, db *sqlx.DB) *AuthMiddleware {
	return &AuthMiddleware{
		client: client,
		db:     db,
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
			log.Print(err.Error())
			http.Error(w, "Failed to verify token", http.StatusForbidden)
			return
		}
		user, err := auth.client.GetUser(r.Context(), token.UID)

		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			return
		}

		u := toUser(user)
		_, syncErr := repository.SyncUser(auth.db, &u)
		if syncErr != nil {
			log.Print(syncErr.Error())
			http.Error(w, "Failed to sync user", http.StatusInternalServerError)
			return
		}

		ctx := httputil.SetUserToContext(r.Context(), &u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenFromHeader(req *http.Request) (string, error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header not found")
	}

	l := len(bearer)
	if len(header) > l+1 && header[:l] == bearer {
		return header[l+1:], nil
	}

	return "", errors.New("authorization header format must be 'Bearer {token}'")
}

func toUser(u *auth.UserRecord) model.User {
	return model.User{
		FirebaseUID: u.UID,
		Email:       u.Email,
		PhotoURL:    u.PhotoURL,
		DisplayName: u.DisplayName,
	}
}
