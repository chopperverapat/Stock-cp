package api

import (
	"server/db"

	"github.com/gin-gonic/gin"
)

// setup call func to setup route
func Setup(router *gin.Engine) {

	db.SetupDB()
	db.SetupDBlog()
	SetupAuthenApi(router)
	SetupProductApi(router)
	SetupTransactionApi(router)
	SetupLog(router)
}
