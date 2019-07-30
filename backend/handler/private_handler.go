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
	user := model.GetUserFromContext(r)
	comments, err := model.GetComments(h.dbx, user.FirebaseUID)
	if err != nil {
		log.Print(err)
		util.WriteJSON(nil, w, http.StatusInternalServerError)
		return
	}
	resp := util.Response{
		Message: fmt.Sprintf(" Hello %s from private endpoint! You have %d comments", user.Email, len(comments)),
	}
	util.WriteJSON(resp, w, http.StatusOK)
}
