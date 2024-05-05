package req

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func New(c *gin.Context) Context {
	id := uuid.New().String()
	c.Header("X-Request-ID", string(id))
	c.Set("request_id", string(id))

	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.Scope().SetTag("request-id", string(id))
	}

	// TODO: actually parse headers, params, and body
	parsed := Parsed{
		Headers: make(map[string]string),
		Params:  make(map[string]string),
		Body:    []byte{},
	}

	return Context{
		ID:      id,
		Parsed:  parsed,
		Context: c,
	}
}
