package handler

import (
	"fmt"

	"github.com/simiancreative/services/server/req"
	"github.com/simiancreative/services/service"
	"github.com/simiancreative/services/service/errors"
)

func auth(
	req req.Context,
	config service.Config,
) error {
	var err error

	if config.Auth == nil {
		return nil
	}

	err = config.Auth(req)
	if err != nil {
		err = fmt.Errorf("Authentication Failed: %v", err.Error())
		resultErr := errors.Common.Result(errors.NotAuthorized, err.Error())
		return resultErr
	}

	return nil
}
