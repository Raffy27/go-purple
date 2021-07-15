package middleware

import (
	"strings"
	"time"

	"github.com/Raffy27/go-purple/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authentication")
	if authHeader == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Authentication header is missing"})
		return
	}

	tmp := strings.Split(authHeader, "Bearer")
	if len(tmp) < 2 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid token"})
		return
	}
	tokenString := strings.TrimSpace(tmp[1])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := config.Get().GetString("jwt.secret")
		return []byte(secretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !(ok && token.Valid) {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}
	// Check for an expired token
	exp := time.Unix(int64(claims["expires"].(float64)), 0)
	if time.Now().After(exp) {
		c.AbortWithStatusJSON(401, gin.H{"error": "Token has expired"})
		return
	}

	c.Set("claims", claims)
	c.Next()
}
