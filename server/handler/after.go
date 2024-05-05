package handler

import (
	"github.com/simiancreative/services/server/req"
	"github.com/simiancreative/services/service"
)

func after(config service.Config, req req.Context) {
	if config.After == nil {
		return
	}

	config.After(req)
}
