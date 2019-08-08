package sample

import (
	"net/http"
)

type PublicHandler struct{}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

func (h *PublicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "Hello from a public endpoint! You don't need to be authenticated to see this.",
	}
	WriteJSON(resp, w, http.StatusOK)
}
