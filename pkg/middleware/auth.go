package middleware

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware(c *gin.Context) {
	_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SIGNINGKEY"))
		return b, nil
	})

	if err == nil {
		c.Next()
	} else {
		c.JSON(401, gin.H{"error": fmt.Sprint(err)})
		c.Abort()
	}
}
