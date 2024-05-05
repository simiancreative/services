package server

import (
	"github.com/gin-gonic/gin"

	"github.com/simiancreative/services/server/middlewares"
)

var router *gin.Engine

func Init() {
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()

	router.Use(
		middlewares.Logger,
		middlewares.Recovery,
		middlewares.Errors,
	)
}

func GetRouter() *gin.Engine {
	return router
}
