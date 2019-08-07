package handler

import (
	"net/http"

	"github.com/voyagegroup/treasure-app/util"
)

type PublicHandler struct{}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

func (h *PublicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := util.Response{
		Message: "Hello from a public endpoint! You don't need to be authenticated to see this.",
	}
	util.WriteJSON(resp, w, http.StatusOK)
}
