package users

import (
	"fmt"

	"github.com/jsuman/kennhouse-users-api/src/datasource/mysql/usersdb"
	"github.com/jsuman/kennhouse-users-api/src/logger"
	"github.com/jsuman/kennhouse-users-api/src/utils/cryptoUtils"
	datetimeutils "github.com/jsuman/kennhouse-users-api/src/utils/dateTimeUtils"
	"github.com/jsuman/kennhouse-users-api/src/utils/errors"
)

func (user *User) Get() *errors.RestErr {
	if connectionErr := usersdb.Client.Ping(); connectionErr != nil {
		panic(connectionErr)
	}
	stmt, err := usersdb.Client.Prepare("select id, first_name, last_name, email, date_created, status from users_db.users where id = ?;")
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while trying to get the user by Id", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare("insert into users_db.users (first_name, last_name, email, date_created, status, password) values (?, ?, ?, ?, ?, ?);")
	if err != nil {
		logger.Error("error while trying to prepare save statement ", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	user.Status = StatusActive
	user.DateCreated = datetimeutils.GetNowDbString()
	pass, passErr := cryptoUtils.EncryptPassword(user.Password)
	if passErr != nil {
		return errors.InternalServerError(err.Error())
	}
	user.Password = pass
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error while trying to insert the user ", err)
		return errors.InternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while trying to get the last inserted user id ", err)
		return errors.InternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare("update users_db.users set first_name=?, last_name=?, email=? where id=?;")
	if err != nil {
		logger.Error("error while trying to prepare the user update statement", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error while trying to update the user ", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() (bool, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare("delete from users where id = ?;")
	if err != nil {
		logger.Error("error while trying to prepare the delete user statement", err)
		return false, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error while trying to delete the user", err)
		return false, errors.InternalServerError("database error")
	}
	return true, nil
}

func (user *User) FindUser(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare("select id, first_name, last_name, email, date_created, status from users where status =?;")
	if err != nil {
		logger.Error("error while trying to prepare the find statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error while trying to find the user ", err)
		return nil, errors.InternalServerError("database error")
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error while trying to scan the user row into the user struct", err)
			return nil, errors.InternalServerError("database error")

		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users found with the status %s ", status))
	}
	return users, nil
}
