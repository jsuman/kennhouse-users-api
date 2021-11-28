package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jsuman/kennhouse-users-api/logger"
)

var router = gin.Default()

func StartApplication() {
	logger.Info("about to start the application..")
	mapUrls()
	router.Run(":8080")
}
