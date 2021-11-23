package users

import (
	"fmt"

	"github.com/jsuman/kennhouse-users-api/datasource/mysql/usersdb"
	datetimeutils "github.com/jsuman/kennhouse-users-api/utils/dateTimeUtils"
	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

func (user *User) Get() *errors.RestErr {
	if connectionErr := usersdb.Client.Ping(); connectionErr != nil {
		panic(connectionErr)
	}
	stmt, err := usersdb.Client.Prepare("select id, first_name, last_name, email, date_created from users_db.users where id = ?;")
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.InternalServerError(fmt.Sprintf("error while trying to get userId %d . %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare("insert into users_db.users (first_name, last_name, email, date_created) values (?, ?, ?,?);")
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, datetimeutils.GetNow())
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error occured while saving the user %s ", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error occured while saving the user %s ", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare("update users_db.users set first_name=?, last_name=?, email=? where id=?;")
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error occured while updating the user %s ", err.Error()))
	}
	return nil
}
