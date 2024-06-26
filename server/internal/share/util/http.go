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
	// Set Content-Type before writing
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func BuildErrResponse(message string) func(error) map[string]string {
	return func(err error) map[string]string {
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		return map[string]string{
			"message": message,
			"error":   errStr,
		}
	}
}

func BuildResponse(message string) map[string]string {
	return map[string]string{
		"message": message,
	}
}
