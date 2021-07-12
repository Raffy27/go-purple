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
	r := NewRouter()

	srv = &http.Server{
		Addr:    config.Get().GetString("server.address"),
		Handler: r,
	}
	go srv.ListenAndServe()

	log.Println("Server initialized")

	return r
}

func Shutdown() {
	sec := time.Duration(config.Get().GetInt("server.maxShutdownSec")) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), sec)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Shutdown error:", err)
		return
	}
	log.Println("Server stopped")
}
