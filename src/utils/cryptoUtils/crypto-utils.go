package cryptoUtils

import (
	"github.com/jsuman/kennhouse-users-api/src/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *errors.RestErr) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.InternalServerError("internal server error")
	}
	return string(hashedPassword), nil
}

func ComparePassword(storedPassword string, credentialPassword string) *errors.RestErr {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentialPassword)); err != nil {
		return errors.InternalServerError("password doesn't match")
	}
	return nil
}
