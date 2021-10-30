package app

import (
	"github.com/gin-gonic/gin"
	"repository.com/my_username/repo_name/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("*** Begin start appication ***")
	router.Run(":8081")
}
