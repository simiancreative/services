package handler

import (
	"github.com/simiancreative/services/server/req"
	"github.com/simiancreative/services/service"
)

func before(config service.Config, req req.Context) error {
	if config.Before == nil {
		return nil
	}

	return config.Before(req)
}
