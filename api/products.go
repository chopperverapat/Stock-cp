package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"server/db"
	"server/logger"
	"server/model"
)

func SetupProductApi(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// /product?token=1234
		productAPI.GET("/product", getproduct)
		// productAPI.GET("/product", interceptor.JwtVerify, getproduct)
		// productAPI.GET("/product/:id", interceptor.JwtVerify, getproductbyID)
		productAPI.GET("/product/:id", getproductbyID)
		productAPI.POST("/product", createproduct)
		// productAPI.POST("/product", interceptor.JwtVerify, createproduct)
		// productAPI.PUT("/product", interceptor.JwtVerify, updateproduct)
		productAPI.PUT("/product", updateproduct)

	}
}

func getproduct(c *gin.Context) {
	var productKeyword []model.Product
	// get product if have keyword search
	keyword := c.Query("keyword")
	if keyword == "" {
		db.GetDB().Find(&productKeyword)
		logger.Savelog(c)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"resutl": gin.H{
				"count": (len(productKeyword)),
				"data":  productKeyword}})
	} else {

		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("product_name like ?", keyword).Find(&productKeyword)
		logger.Savelog(c)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"resutl": gin.H{
				"count": (len(productKeyword)),
				"data":  productKeyword}})
		logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)

	}
}

func getproductbyID(c *gin.Context) {
	var productID model.Product
	result := db.GetDB().Where("id=?", c.Param("id")).First(&productID)
	if result.Error != nil {
		logger.Savelog(c)
		c.JSON(http.StatusNotFound, gin.H{"status": "not ok", "error": result.Error})
		logger.ErrorLogger.Printf("An error occurred: %v", result.Error)

		return
	}
	logger.Savelog(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": productID})
	logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)

}

func createproduct(c *gin.Context) {
	var product model.Product

	product.ProductName = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64) // return ค่า,err ไม่ได้ใช้ เลยใช้ _
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)   // return ค่า,err
	product.CreatedAt = time.Now()

	db.GetDB().Create(&product)
	saveImagetoDB(c, &product)
	logger.Savelog(c)
	c.JSON(http.StatusOK, gin.H{"status": "created already", "result": product})
	logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)

}

func saveImagetoDB(c *gin.Context, product *model.Product) {

	image, err := c.FormFile("image")
	if err != nil {
		return
	}
	if image != nil {
		product.Image = image.Filename
		currentDir, _ := os.Getwd()
		extension := filepath.Ext(image.Filename)
		filename := fmt.Sprintf("product-prod_%d%s", product.ID, extension)
		filepath := fmt.Sprintf("%s/uploaded/images-prod/%s", currentDir, filename)
		if fileExists(filepath) {
			os.Remove(filepath)
		}
		logger.Savelog(c)
		c.SaveUploadedFile(image, filepath)
		db.GetDB().Model(&product).Update("image", filename)

	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func updateproduct(c *gin.Context) {
	var editProduct model.Product
	// convert str to int ,base 10 , int32
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	editProduct.ID = uint(id)
	editProduct.ProductName = c.PostForm("name")
	editProduct.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	editProduct.Price, _ = strconv.ParseFloat(c.PostForm("stock"), 64)

	// Log the values of editProduct for debugging purposes
	fmt.Printf("editProduct: %v\n", editProduct)
	var checkDB model.Product
	result := db.GetDB().First(&checkDB, editProduct.ID)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{"status": result.Error})
		logger.ErrorLogger.Printf("An error occurred: %v", result.Error)
		return
	} else {
		editProduct.Image = checkDB.Image
		// save image
		db.GetDB().Save(&editProduct)
		saveImagetoDB(c, &editProduct)
		logger.Savelog(c)
		c.JSON(http.StatusOK, gin.H{"status": "updated product successfully", "data": editProduct})
		logger.InfoLogger.Printf("API Endpoint: %s, Request Method: %s", c.FullPath(), c.Request.Method)
	}
}
