package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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
		return nil, err
	}

	return newUser, nil
}
