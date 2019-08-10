package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/voyagegroup/treasure-app/httputil"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)
	if err != nil {
		respondErrorJson(w, status, err)
		return
	}
	respondJSON(w, status, res)
	return
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondErrorJson(w http.ResponseWriter, code int, err error) {
	log.Printf("code=%d, err=%s", code, err)
	if e, ok := err.(*httputil.HTTPError); ok {
		respondJSON(w, code, e)
	} else if err != nil {
		he := httputil.HTTPError{
			Message: err.Error(),
		}
		respondJSON(w, code, he)
	}
}
