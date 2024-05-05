package server

import (
	"os"

	"github.com/simiancreative/simiango/logger"

	"github.com/simiancreative/services/server/handler"
	"github.com/simiancreative/services/service"
)

func Start() {
	for _, service := range services {
		initService(service)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Debug(
		"Server: initialized and starting",
		logger.Fields{"port": port},
	)

	if err := router.Run(); err != nil {
		logger.Fatal(
			"Server: failed to start",
			logger.Fields{"err": err},
		)
	}
}

func initService(config service.Config) {
	logger.Debug("Server: adding route", logger.Fields{
		"method": config.Method,
		"path":   config.Path,
	})

	router.Handle(config.Method, config.Path, handler.New(config))
}
