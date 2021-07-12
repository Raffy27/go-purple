package controllers

import (
	"net/http"
	"time"

	"github.com/Raffy27/go-purple/models"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (t *TestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Pong!",
		"time": time.Now(),
	})
}

func (t *TestController) Profile(c *gin.Context) {
	if tmp, ok := c.Get("user"); ok {
		user := tmp.(models.User)
		c.JSON(http.StatusOK, gin.H{
			"msg":       "This is your profile!",
			"user":      user.Username,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid context encountered?",
		})
	}
}
