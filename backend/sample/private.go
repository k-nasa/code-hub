package sample

import (
	"fmt"
	"log"
	"net/http"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/repository"

	"github.com/jmoiron/sqlx"
)

type PrivateHandler struct {
	db *sqlx.DB
}

func NewPrivateHandler(db *sqlx.DB) *PrivateHandler {
	return &PrivateHandler{
		db: db,
	}
}

func (h *PrivateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contextUser, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		log.Print(err)
		WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	user, err := repository.GetUser(h.db, contextUser.FirebaseUID)
	if err != nil {
		log.Printf("Show user failed: %s", err)
		WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	resp := Response{
		Message: fmt.Sprintf("Hello %s from private endpoint! Your firebase uuid is %s", user.DisplayName, user.FirebaseUID),
	}
	WriteJSON(resp, w, http.StatusOK)
}
