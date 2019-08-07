package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/util"
	"log"
	"net/http"
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
	user, err := model.GetUser(h.dbx, model.GetUserFromContext(r).FirebaseUID)
	if err != nil {
		log.Printf("Get user failed: %s", err)
		util.WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	resp := util.Response{
		Message: fmt.Sprintf("Hello %s from private endpoint! Your firebase uuid is %s", user.DisplayName, user.FirebaseUID),
	}
	util.WriteJSON(resp, w, http.StatusOK)
}
