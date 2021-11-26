package services

import (
	"github.com/jsuman/kennhouse-users-api/domain/users"
	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func SearchUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteUser(userId int64) (bool, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return false, err
	}
	qResult, err := result.Delete()
	if err != nil {
		return false, err
	}
	return qResult, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := SearchUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.Email != "" {
			currentUser.Email = user.Email
		}
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

	} else {
		currentUser.Email = user.Email
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}
