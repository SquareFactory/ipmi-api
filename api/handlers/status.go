package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/squarefactory/ipmitool"
)

func Status(c *gin.Context) {

	hostname := c.Param("host")
	hostIP := os.Getenv(hostname)

	if len(hostIP) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"error": "Host not defined"})
		return
	}

	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	cl, err := ipmitool.NewClient(hostIP, 0, username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status, err := cl.Power.Status()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
