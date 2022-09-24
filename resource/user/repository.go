package user

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/qna-page/qna-page/utils"
)

type UserRepository interface {
	List() (*[]User, error)
	FindByID(id string) (*User, error)
	Create(email, displayName, password string) (*User, error)
}

// Models what we inject into the repo
type UserRepo struct {
	db *sqlx.DB
}

func InitRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) List() (*[]User, error) {
	var users []User
	err := r.db.Select(&users, "SELECT * FROM user")
	if err != nil {
		fmt.Println("Error", err)
	}
	return &users, nil
}

func (r *UserRepo) FindByID(id string) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &user, nil
}

func (r *UserRepo) Create(email, displayName, password string) (*User, error) {
	newUser := &User{}
	newUser.GenerateId()
	newUser.Email = email
	newUser.DisplayName = displayName
	newUser.HashPassword(password)

	_, err := r.db.NamedExec("INSERT INTO user (id, email, hash, display_name) VALUES (:id, :email, :hash, :display_name);", *newUser)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "user.email") {
			message := &utils.DBFieldError{Detail: &UserInputJSON{Email: "A user with this email already exists."}}
			return nil, message
		}

		if strings.Contains(msg, "user.display_name") {
			message := &utils.DBFieldError{Detail: &UserInputJSON{DisplayName: "A user with this display name already exists."}}
			return nil, message
		}

		// Avoid exposing unhandled errors to end-user
		return nil, &utils.MaskError{Err: msg}
	}

	return newUser, nil
}
