package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}

func ReturnHTTPError(w http.ResponseWriter, code int, e *ErrorMessage) {
	response := &ErrorMessage{}

	// Use the default text code if a title is not set
	if e.Title == "" {
		response.Title = http.StatusText(code)
	}

	// Serialize to JSON
	b, _ := json.Marshal(&e)

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)

	return
}
