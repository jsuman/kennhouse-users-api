package services

import (
	"github.com/jsuman/kennhouse-users-api/domain/users"
	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

var (
	UserService UserServiceInterface = &userService{}
)

type userService struct {
}

type UserServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	SearchUser(int64) (*users.User, *errors.RestErr)
	DeleteUser(int64) (bool, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	FindUser(string) (users.Users, *errors.RestErr)
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) SearchUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) DeleteUser(userId int64) (bool, *errors.RestErr) {
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

func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := u.SearchUser(user.Id)
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

func (u *userService) FindUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindUser(status)
}
