package features

import (
	"fmt"
	"os"

	"github.com/mandrigin/gin-spa/spa"
	"github.com/simiancreative/simiango/logger"

	"github.com/simiancreative/services/server"
)

func SPA(url string, dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		panic(fmt.Sprintf("spa dir does not exist: %v", dir))
	}

	logger.Debug("Server: adding spa route", logger.Fields{
		"path": url,
		"dir":  dir,
	})

	server.GetRouter().Use(spa.Middleware(url, dir))
}
