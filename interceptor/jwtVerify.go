package interceptor

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"server/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret string = "123455678"

func JwtSign(userDB model.User) (string, error) {
	jwtGOon := jwt.MapClaims{}
	jwtGOon["id"] = userDB.ID
	jwtGOon["username"] = userDB.Username
	jwtGOon["level"] = userDB.Level
	jwtGOon["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtGOon)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtVerify(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		c.Abort()
		return
	}

	if len(strings.Split(authHeader, " ")) < 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid header"})
		c.Abort()
		return
	}
	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(secret), nil

	})

	if jwtMap, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(jwtMap["id"])

		staffID := fmt.Sprintf("%v", jwtMap["id"])
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", jwtMap["username"])
		c.Set("jwt_level", jwtMap["level"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid token", "error": err})
		c.Abort()
	}
}
