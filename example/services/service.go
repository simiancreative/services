package services

import (
	"github.com/simiancreative/services/server/req"
	"github.com/simiancreative/services/service"
)

var Config = &service.Config{
	Method:  "GET",
	Path:    "/services",
	Handler: handler,
}

func handler(ctx req.Context) (interface{}, error) {
	return "Hello, World!", nil
}
