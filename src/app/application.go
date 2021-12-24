package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jsuman/kennhouse-users-api/src/logger"
)

var router = gin.Default()

func StartApplication() {
	logger.Info("about to start the application..")
	mapUrls()
	router.Run(":8089")
}
