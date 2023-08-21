package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InterceptorGeneral(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid token"})
		c.Abort()
	}
}
