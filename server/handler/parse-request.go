package handler

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/simiancreative/services/server/req"
)

func parseRequest(c *gin.Context) req.Context {
	id := uuid.New().String()
	c.Header("X-Request-ID", id)
	c.Set("request_id", id)

	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.Scope().SetTag("request-id", string(id))
	}

	return req.Context{
		ID:      id,
		Context: c,
	}
}
