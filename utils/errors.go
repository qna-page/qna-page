package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Use to obscure unhandled errors
type MaskError struct {
	Err string
}

func (e *MaskError) Error() string {
	fmt.Printf("Unhandled Error: %v", e.Err)
	return ""
}

type DBFieldError struct {
	Detail interface{}
}

func (e *DBFieldError) Error() string {
	return "A DB error occurred for a request field."
}

type ErrorMessage[D any] struct {
	Title  string `json:"title"`
	Detail D      `json:"detail,omitempty"`
}

type UnmarshalMessage struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func HTTPError(w http.ResponseWriter, code int, e error) {
	var response any
	title := http.StatusText(code)

	switch err := e.(type) {
	case *MaskError:
		// Use this error type for all "unhandled" errors
		_ = err.Error()
		code = 500
		response = &ErrorMessage[string]{Title: "Internal Server Error"}
	case *json.UnmarshalTypeError:
		detail := &UnmarshalMessage{err.Field, fmt.Sprintf("%v", err.Type), err.Value}
		response = &ErrorMessage[*UnmarshalMessage]{Title: title, Detail: detail}
	case *json.SyntaxError:
		response = &ErrorMessage[string]{Title: title, Detail: "Invalid JSON syntax"}
	case validation.Errors:
		response = &ErrorMessage[error]{Title: title, Detail: e}
	case *DBFieldError:
		response = &ErrorMessage[interface{}]{Title: title, Detail: err.Detail}
	case error:
		response = &ErrorMessage[string]{Title: title, Detail: err.Error()}
	default:
		response = &ErrorMessage[string]{Title: title, Detail: "Unhandled Application Error"}
		fmt.Printf("Unhandled Error: %s\n", reflect.TypeOf(e))
	}

	// Serialize to JSON
	b, _ := json.Marshal(response)

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
