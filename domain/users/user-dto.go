package users

import (
	"net/mail"

	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.BadRequestError("invalid email address")
	}
	return nil
}
