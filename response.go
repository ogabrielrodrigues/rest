package rest

import (
	"encoding/json"
	"net/http"
)

// Write http status code and return data in JSON format
func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(data)
}

// Write http status code and return empty
func End(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
