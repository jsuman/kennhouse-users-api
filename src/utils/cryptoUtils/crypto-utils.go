package cryptoUtils

import (
	"encoding/hex"

	"github.com/jsuman/kennhouse-users-api/src/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *errors.RestErr) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.InternalServerError("internal server error")
	}
	return hex.EncodeToString(hashedPassword), nil
}
