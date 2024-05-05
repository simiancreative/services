package features

import (
	"github.com/gin-gonic/gin"

	"github.com/simiancreative/services/server"
)

func EnableHealthCheck() {
	router := server.GetRouter()
	router.Handle("GET", "/status", healthGET)
}

func healthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
