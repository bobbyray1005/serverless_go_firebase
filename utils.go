package function

import (
	"encoding/json"
	"net/http"
)

func respond(httpStatus int, responseJSON interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(responseJSON)
}

// responds just with a status
func respondStatus(httpStatus int, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
}
