package app

import (
	"github.com/jsuman/kennhouse-users-api/src/controllers/ping"
	"github.com/jsuman/kennhouse-users-api/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.FindUser)
	router.POST("/users/login", users.LoginUser)
}
