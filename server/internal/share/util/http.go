package util

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Handle error (e.g., log the error and return a bad request status code)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}