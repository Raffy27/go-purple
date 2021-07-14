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
	tmp, _ := c.Get("user")
	user := tmp.(*models.User)
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"user":      user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
}
