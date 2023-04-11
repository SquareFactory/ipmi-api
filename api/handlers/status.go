package handlers

import (
	"net/http"

	"github.com/Bancadati/ipmitool"
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {

	hostname := c.Param("host")
	cl, err := ipmitool.NewClient(hostname, 0, "IPMIUSER", "Password")
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
