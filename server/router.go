package server

import (
	"github.com/Raffy27/go-purple/controllers"
	"github.com/Raffy27/go-purple/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Static files
	r.Use(static.Serve("/", static.LocalFile("public", false)))

	test := new(controllers.TestController)
	r.GET("/ping", test.Ping)

	auth := new(controllers.Auth)
	r.POST("/login", auth.Login)
	r.POST("/create", auth.Create)

	safe := r.Group("api")
	{
		safe.Use(middleware.Authentication)
		safe.GET("/profile", test.Profile)
	}

	return r
}
