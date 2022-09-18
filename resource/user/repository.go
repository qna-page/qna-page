package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	List() (*[]User, error)
	FindByID(Id string) (*User, error)
	Save(user *User) error
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

func (r *UserRepo) FindByID(Id string) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM user WHERE id=$1", Id)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &user, nil
}

// Save - everything is skeleton'ed for now
func (r *UserRepo) Save(user *User) error {
	// TODO: implement bcrypt for passwords
	return nil
}
