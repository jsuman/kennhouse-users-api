package mysqlUtils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jsuman/kennhouse-users-api/src/utils/errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	fmt.Println(err.Error())
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NotFoundError("No matching record found")
		}
		return errors.InternalServerError("Error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("Invalid data")
	}
	return errors.InternalServerError("Error processing request")
}
