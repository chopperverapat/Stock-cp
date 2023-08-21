package api

import (
	"fmt"
	"net/http"
	"time"

	"server/db"
	"server/interceptor"
	"server/logger"
	"server/model"

	"github.com/gin-gonic/gin"
)

func SetupTransactionApi(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/transaction", gettransaction)
		productAPI.POST("/transaction", interceptor.JwtVerify, createtransaction)
	}
}

type joinTransactionUSer struct {
	ID            uint
	Total         float64
	Paid          float64
	Change        float64
	PaymentType   string
	PaymentDetail string
	OrderList     string
	Staff         string
	CreatedAt     time.Time
}

func gettransaction(c *gin.Context) {
	var result []joinTransactionUSer
	query := "SELECT transactions.id, total, paid, change, payment_type, payment_detail, " +
		"order_list, users.username as staff, transactions.created_at " +
		"FROM transactions " +
		"JOIN users ON transactions.staff_id = users.id"

	db.GetDB().Debug().Raw(query, nil).Scan(&result)
	logger.Savelog(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "count": len(result), "data": result})
	logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)
}

func createtransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		// fmt.Println(transaction)
		transaction.StaffID = c.GetString("jwt_staff_id")
		fmt.Printf("ID : %v\n", transaction.StaffID)
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		logger.Savelog(c)
		c.JSON(http.StatusOK, gin.H{"status": "created transaction", "data": transaction})
		logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)
	} else {
		logger.Savelog(c)
		c.JSON(http.StatusNotFound, gin.H{"status": "not ok", "error": err})
		logger.ErrorLogger.Printf("An error occurred: %v", err)
	}
}
