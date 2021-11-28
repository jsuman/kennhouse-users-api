package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jsuman/kennhouse-users-api/domain/users"
	"github.com/jsuman/kennhouse-users-api/services"
	"github.com/jsuman/kennhouse-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, updateErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if updateErr != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err.Message)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, uErr := services.UserService.UpdateUser(isPartial, user)
	if uErr != nil {
		c.JSON(uErr.Status, uErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func SearchUser(c *gin.Context) {
	userId, searchErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if searchErr != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err.Message)
		return
	}
	user, getErr := services.UserService.SearchUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, deleteErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if deleteErr != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err.Message)
		return
	}
	status, getErr := services.UserService.DeleteUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, status)
}

func FindUser(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.FindUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
