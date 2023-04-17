package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/squarefactory/ipmi-api/api/handlers"
)

func main() {

	r := gin.Default()

	r.POST("/host/:host/on", handlers.PowerOn)
	r.POST("/host/:host/off", handlers.PowerOff)
	r.POST("/host/:host/cycle", handlers.Cycle)
	r.GET("/host/:host/status", handlers.Status)
	r.POST("/host/:host/soft", handlers.Soft)
	r.POST("/host/:host/reset", handlers.Reset)

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if len(listenAddress) == 0 {
		listenAddress = ":8080"
	}

	err := r.Run(listenAddress)
	if err != nil {
		log.Fatal(err)
	}
}
