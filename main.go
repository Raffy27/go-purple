package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Raffy27/go-purple/config"
	"github.com/Raffy27/go-purple/server"
	"github.com/Raffy27/go-purple/server/db"
)

func main() {

	config.Init()
	db.Init()
	server.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	log.Println("Stopping")

	server.Shutdown()
	db.Shutdown()
}
