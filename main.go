package main

import (
	"log"

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
	//r.POST("/host/{host}/boot", handlers.boot)
	//r.GET("/host/{host}/stat", handlers.stat)
	//r.GET("/host/{host}/state", handlers.state)
	//r.POST("/host/{host}/wake", handlers.wake)
	//r.POST("/host/{host}/suspend", handlers.suspend)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
