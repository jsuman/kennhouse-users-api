package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type Users []User

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	jsonuser, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(jsonuser, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(jsonuser, &privateUser)
	return privateUser
}
