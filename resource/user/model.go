package user

import (
	"github.com/jaevor/go-nanoid"
	"golang.org/x/crypto/bcrypt"
)

var UserSchema = `
CREATE TABLE IF NOT EXISTS user (
	id				TEXT	PRIMARY KEY,
	email			TEXT	NOT NULL,
	hash			TEXT	NOT NULL,
	display_name	TEXT	NOT NULL,
	is_super_admin	INTEGER DEFAULT 0
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email 
ON user (email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_display_name 
ON user (display_name);
`

type User struct {
	Id           string `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Hash         string `db:"hash" json:"-"`
	DisplayName  string `db:"display_name" json:"displayName"`
	IsSuperAdmin bool   `db:"is_super_admin" json:"isSuperAdmin"`
}

func (u *User) GenerateId() {
	nanoId, err := nanoid.Standard(16)
	if err != nil {
		panic(err)
	}
	u.Id = nanoId()
}

func (u *User) HashPassword(password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	u.Hash = string(bytes)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}
