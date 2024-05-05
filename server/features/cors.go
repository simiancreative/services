package features

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/simiancreative/services/server"
)

func CORS() {
	origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	methods := strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	headers := strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")

	config := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowHeaders:     headers,
		ExposeHeaders:    headers,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	server.GetRouter().Use(cors.New(config))
}
