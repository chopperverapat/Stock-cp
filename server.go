package main

import (
	"server/api"
	"server/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLoggers()

	router := gin.Default()

	router.Static("/images", "./uploaded/images/")

	api.Setup(router)

	router.Run(":8082")
}
