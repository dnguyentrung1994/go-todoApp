package controllers

import (
	"go-todoApp/libs"
	"go-todoApp/services"
)

type JWTAuthController struct {
	logger      libs.Logger
	service     services.JWTAuthService
	userService services.UserService
}

//generates new instance of JWTAuthController
func NewJWTAuthController(logger libs.Logger, userService services.UserService, JWTAuthService services.JWTAuthService) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     JWTAuthService,
		userService: userService,
	}
}
