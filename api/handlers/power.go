package handlers

import (
	"net/http"

	"github.com/Bancadati/ipmitool"
	"github.com/gin-gonic/gin"
)

func PowerOn(c *gin.Context) {
	cl, err := ipmitool.NewClient("192.198.1.1", 0, "IPMIUSER", "Password")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status, err := cl.Power.Status()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if status == ipmitool.PowerStateOff {
		err := cl.Power.On()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Compute node already powered on"})
	}
}

func PowerOff(c *gin.Context) {

	cl, err := ipmitool.NewClient("192.198.1.1", 0, "IPMIUSER", "Password")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status, err := cl.Power.Status()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if status == ipmitool.PowerStateOn {
		err := cl.Power.Off()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Compute node already powered off"})
	}
}

func Cycle(c *gin.Context) {

	cl, err := ipmitool.NewClient("192.198.1.1", 0, "IPMIUSER", "Password")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status, err := cl.Power.Status()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status == ipmitool.PowerStateOn {
		err := cl.Power.Cycle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success"})

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Compute node already powered off"})
	}

}
