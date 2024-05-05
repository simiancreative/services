package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simiancreative/simiango/monitoring/sentry"

	"github.com/simiancreative/services/service/errors"
)

func handleError(err error, c *gin.Context) bool {
	if err == nil {
		return false
	}

	resultErr := convertError(err)
	status := resultErr.Status

	if status < 500 {
		return false
	}

	sentry.GinCaptureError(c, err)

	c.JSON(status, resultErr)

	return true
}

func convertError(err error) *errors.Result {
	resultErr, ok := err.(*errors.Result)

	if !ok {
		resultErr = errors.Common.Result(errors.InternalServerError, err.Error())
		return resultErr
	}

	return resultErr
}
