package features

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/simiancreative/services/server"
)

func AddSentry(scopeFunc func(*gin.Context, *sentry.Scope)) {
	router := server.GetRouter()

	router.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	router.Use(sentrySetup(scopeFunc))
}

func sentrySetup(scopeFunc func(*gin.Context, *sentry.Scope)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sentryScopeFunc", scopeFunc)
		c.Next()
	}
}
