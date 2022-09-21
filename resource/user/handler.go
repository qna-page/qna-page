package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qna-page/qna-page/utils"
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

type UserInputJSON struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

type UserOutputJSON struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

// Create a user from json, returns json
func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	data := &UserInputJSON{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.ReturnHTTPError(w, http.StatusBadRequest, &utils.ErrorMessage{Detail: "Bad Body"})
		return
	}

	user, err := h.userRepo.Create(data.Email, data.DisplayName, data.Password)
	if err != nil {
		utils.ReturnHTTPError(w, http.StatusBadRequest, &utils.ErrorMessage{Detail: "Bad Body"})
		return
	}
	resp := &UserOutputJSON{user.Id, user.Email, user.DisplayName}

	// Cast to JSON
	b, _ := json.Marshal(resp)

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
