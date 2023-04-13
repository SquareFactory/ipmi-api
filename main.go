package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/squarefactory/ipmi-api/api/handlers"
	"github.com/squarefactory/ipmi-api/api/secrets"
)

func main() {

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		secrets.IpmiUsername: secrets.IpmiPassword,
	}))

	authorized.POST("/host/:host/on", handlers.PowerOn)
	authorized.POST("/host/:host/off", handlers.PowerOff)
	authorized.POST("/host/:host/cycle", handlers.Cycle)
	authorized.GET("/host/:host/status", handlers.Status)
	authorized.POST("/host/:host/soft", handlers.Soft)
	authorized.POST("/host/:host/reset", handlers.Reset)
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
