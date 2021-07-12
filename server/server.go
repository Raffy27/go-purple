package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Raffy27/go-purple/config"
	"github.com/gin-gonic/gin"
)

var srv *http.Server

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hewwo")
	})

	srv = &http.Server{
		Addr:    config.Get().GetString("server.address"),
		Handler: r,
	}
	go srv.ListenAndServe()

	log.Println("server initialized")

	return r
}

func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("shutdown error:", err)
	}
}
