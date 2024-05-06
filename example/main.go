package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/simiancreative/simiango/logger"

	"github.com/simiancreative/services/server"
	"github.com/simiancreative/services/server/features"
)

func main() {
	logger.Enable()
	server.Init()

	features.CORS()
	features.AddSentry(sentryScope)
	features.EnableHealthCheck()

	server.Start()
}

func sentryScope(c *gin.Context, s *sentry.Scope) {
	s.SetExtra("extra", "extra")
}
