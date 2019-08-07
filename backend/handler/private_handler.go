package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/domain/model"
	"github.com/voyagegroup/treasure-app/util"
)

type PrivateHandler struct {
	dbx *sqlx.DB
}

func NewPrivateHandler(dbx *sqlx.DB) *PrivateHandler {
	return &PrivateHandler{
		dbx: dbx,
	}
}

func (h *PrivateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contextUser, err := model.GetUserFromContext(r.Context())
	if err != nil {
		log.Print(err)
		util.WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	user, err := model.GetUser(h.dbx, contextUser.FirebaseUID)
	if err != nil {
		log.Printf("Show user failed: %s", err)
		util.WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	resp := util.Response{
		Message: fmt.Sprintf("Hello %s from private endpoint! Your firebase uuid is %s", user.DisplayName, user.FirebaseUID),
	}
	util.WriteJSON(resp, w, http.StatusOK)
}
