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

	v1 := r.Group("api/v1")
	{
		auth := v1.Group("auth")
		{
			c := new(controllers.AuthController)
			auth.POST("login", c.Login)
			auth.POST("logout", c.Logout)
			auth.POST("register", c.Register)
		}

		users := v1.Group("users")
		{
			users.Use(middleware.Authentication())

			c := new(controllers.UserController)
			users.GET("", c.GetAll)
			users.GET("/:user", c.GetByUsername)
		}
	}

	return r
}
