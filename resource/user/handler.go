package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// UserHandler is used to inject everything the handler needs
type UserHandler struct {
	userRepo UserRepository
}

// Initialize a user handler
func InitHandler(userRepo UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// Returns a JSON list of users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: will need pagination
	w.Header().Set("Content-Type", "application/json")

	users, err := h.userRepo.List()
	if err != nil {
		fmt.Println("Error", err)
	}

	b, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// Retrieves a user by ID, returns json
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Look up user in DB
	user, err := h.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error", user)
	}

	// Cast to JSON
	// TODO: We should use a seperate model without password for added safety
	b, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
