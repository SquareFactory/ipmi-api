package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/squarefactory/ipmi-api/api/handlers"
	"github.com/squarefactory/ipmi-api/api/secrets"
)

func main() {

	// set ipmiuser and password
	secrets.ReadSecret()

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

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
