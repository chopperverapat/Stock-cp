package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"server/db"
	"server/interceptor"
	"server/logger"
	"server/model"
)

func SetupAuthenApi(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		var userfromdatabase model.User

		if err := db.GetDB().First(&userfromdatabase, "username = ?", user.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"stutus": "not ok", "error": err})
			logger.Savelog(c)
			logger.ErrorLogger.Printf("An error occurred: %v", err)

		} else {
			equalPass, err := checkPasswordHash(user.Password, userfromdatabase.Password)
			if err != nil {
				logger.Savelog(c)
				c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error", "error": err})
				logger.ErrorLogger.Printf("An error occurred: %v", err)

			} else if !equalPass {
				logger.Savelog(c)
				c.JSON(http.StatusUnauthorized, gin.H{"status": "not ok", "status code": http.StatusUnauthorized})
				logger.ErrorLogger.Printf("An error occurred: %v", err)

			} else {
				token, err := interceptor.JwtSign(userfromdatabase)
				if err != nil {
					logger.Savelog(c)
					c.JSON(http.StatusUnauthorized, gin.H{"status": "cant gen token", "token": token})
					logger.ErrorLogger.Printf("An error occurred: %v", err)

				}
				logger.Savelog(c)
				c.JSON(http.StatusOK, gin.H{"status": "log in", "token": token})
				logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)

			}
		}
	} else {
		logger.Savelog(c)
		c.JSON(http.StatusOK, gin.H{"status": "not ok and cant bind data"})

	}

}

func register(c *gin.Context) {
	var user model.User

	if c.ShouldBind(&user) == nil {
		hashPasswordAleready, err := hashpassword(user.Password)
		if err != nil {
			logger.Savelog(c)
			c.JSON(http.StatusBadRequest, gin.H{"status": "Failed to hash password"})
		}
		user.Password = hashPasswordAleready
		user.CreatedAt = time.Now()
		// check ค่าที่เขียนใน dyabase ว่าถูกต้องไหม ใช้ .Error ดูค่า error ถ่้ามีก็ ขึ้น error
		if err := db.GetDB().Create(&user).Error; err != nil {
			logger.Savelog(c)
			c.JSON(http.StatusOK, gin.H{"status": "not register", "error": err})
			logger.ErrorLogger.Printf("An error occurred: %v", err)
		} else {
			logger.Savelog(c)
			c.JSON(http.StatusOK, gin.H{"status": "register", "data": user})
			logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)

		}
	} else {
		logger.Savelog(c)
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func hashpassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	} else {
		return string(hashedPassword), nil
	}
}

func checkPasswordHash(password, hash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
