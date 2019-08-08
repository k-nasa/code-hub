package sample

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func WriteJSON(v interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, e := w.Write(jsonResponse); e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
