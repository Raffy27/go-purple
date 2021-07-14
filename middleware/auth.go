package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Raffy27/go-purple/config"
	"github.com/Raffy27/go-purple/models"
	"github.com/Raffy27/go-purple/server/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Authentication(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authentication")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Authentication header is missing",
		})
		return
	}

	tmp := strings.Split(authHeader, "Bearer")
	if len(tmp) < 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	tokenString := strings.TrimSpace(tmp[1])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := config.Get().GetString("secrets.jwt")
		return []byte(secretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		username := claims["username"].(string)
		res := db.Main().Collection("users").FindOne(c, gin.H{
			"username": username,
		}, options.MergeFindOneOptions())
		if res.Err() != nil {
			log.Println(res.Err())
			c.AbortWithStatusJSON(402, gin.H{
				"error": "User not found",
			})
			return
		}
		res.Decode(&user)

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Token is not valid",
		})
	}
}
