package users

import (
	"net/mail"
	"strings"

	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.BadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequestError("invalid password")
	}
	return nil
}
