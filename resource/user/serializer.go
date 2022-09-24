package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserInputJSON struct {
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

func (u UserInputJSON) Validate() error {
	return validation.ValidateStruct(&u,
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&u.Email, validation.Required, is.Email),
		// Password must be between 8 and 128 characters
		validation.Field(&u.Password, validation.Required, validation.Length(8, 128)),
		// DisplayName must be between 3 and 32 characters
		validation.Field(&u.DisplayName, validation.Required, validation.Length(3, 32)),
	)
}

type UserOutputJSON struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}
