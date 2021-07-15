package controllers

import (
	"net/http"
	"time"

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
	claims, _ := c.Get("claims")
	c.JSON(http.StatusOK, claims)
}
