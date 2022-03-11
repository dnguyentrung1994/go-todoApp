package routes

import (
	"go-todoApp/api/controllers"
	"go-todoApp/api/middlewares"
	"go-todoApp/libs"
)

type UserRoutes struct {
	logger         libs.Logger
	handler        libs.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
}

func (s UserRoutes) SetUp() {
	s.logger.Info("setting up routes for user-related functionalities...")
	api := s.handler.Gin.Group("/api/user").Use(s.authMiddleware.Handler())
	{
		api.POST("/", s.userController.CreateNewUser)
	}
}

func NewUserRoutes(
	logger libs.Logger,
	handler libs.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		logger:         logger,
		handler:        handler,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
