package server

import (
	"github.com/Raffy27/go-purple/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	test := new(controllers.TestController)
	r.GET("/ping", test.Ping)

	auth := new(controllers.Auth)
	r.POST("/login", auth.Login)
	r.POST("/create", auth.Create)

	return r
}