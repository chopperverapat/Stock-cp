package api

import (
	"net/http"
	"server/db"
	"server/interceptor"
	"server/logger"
	"server/model"

	"github.com/gin-gonic/gin"
)

func SetupLog(router *gin.Engine) {
	logAPI := router.Group("/api/v2")
	{
		// logAPI.GET("/log", getLog)
		logAPI.GET("/log", interceptor.JwtVerify, getLog)
	}
}

func getLog(c *gin.Context) {
	var listlog []model.Log
	db.GetDBlog().Find(&listlog)
	logger.Savelog(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": listlog})

}
