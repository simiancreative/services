package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/simiancreative/services/server/req"
	"github.com/simiancreative/services/service"
)

func New(config service.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		parsedReq := req.New(c)

		if err := auth(parsedReq, config); handleError(err, c) {
			return
		}

		if err := before(config, parsedReq); handleError(err, c) {
			return
		}

		result, err := config.Handler(parsedReq)
		if handleError(err, c) {
			return
		}

		go after(config, parsedReq)

		if result == nil {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
