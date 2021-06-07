package app

import (
	"github.com/gin-gonic/gin"
	"github.com/murilloarturo/bookstore_users_api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8081")
}
