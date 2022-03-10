package middlewares

import (
	"go-todoApp/libs"

	"github.com/gin-contrib/cors"
)

type CorsMiddleware struct {
	handler libs.RequestHandler
	logger  libs.Logger
}

func NewCorsMiddleware(handler libs.RequestHandler, logger libs.Logger) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
	}
}

func (m CorsMiddleware) SetUp() {
	m.logger.Info("Initializing Cors middleware")

	m.handler.Gin.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}))
}
