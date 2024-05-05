package middlewares

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/simiancreative/simiango/logger"
	"github.com/simiancreative/simiango/meta"
	"github.com/simiancreative/simiango/service"
)

func Recovery(c *gin.Context) {
	defer recoverGinPanic(c, handleRecovery)
	c.Next()
}

func recoverGinPanic(c *gin.Context, buildResp func(*gin.Context, map[string]interface{})) {
	if err := recover(); err != nil {
		// Check for a broken connection, as it is not really a
		// condition that warrants a panic stack trace.
		var brokenPipe bool
		if ne, ok := err.(*net.OpError); ok {
			var se *os.SyscallError
			if errors.As(ne, &se) {
				seStr := strings.ToLower(se.Error())
				if strings.Contains(seStr, "broken pipe") ||
					strings.Contains(seStr, "connection reset by peer") {
					brokenPipe = true
				}
			}
		}

		id, _ := c.Get("request_id")
		stack := meta.Stack()
		httpRequest, _ := httputil.DumpRequest(c.Request, false)
		headers := strings.Split(string(httpRequest), "\r\n")
		for idx, header := range headers {
			current := strings.Split(header, ":")
			if current[0] == "Authorization" {
				headers[idx] = current[0] + ": *"
			}
		}

		logger.Error(
			err.(error).Error(),
			logger.Fields{
				"request_id": id,
				"request":    headers,
				"stack":      stack,
			},
		)

		if brokenPipe {
			// If the connection is dead, we can't write a status to it.
			c.Error(err.(error)) //nolint: errcheck
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		buildResp(c, map[string]interface{}{
			"request_id": id,
			"error":      err.(error).Error(),
			"headers":    headers,
			"stack":      stack,
		})
	}
}

func handleRecovery(c *gin.Context, context map[string]interface{}) {
	logLevel := logger.Level()

	if logLevel > logger.Levels[4] {
		c.JSON(http.StatusInternalServerError, service.ResultError{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Reasons: []map[string]interface{}{context},
		})
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}
