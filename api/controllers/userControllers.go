package controllers

import (
	"go-todoApp/constants"
	"go-todoApp/entities"
	"go-todoApp/libs"
	"go-todoApp/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service services.UserService
	logger  libs.Logger
}

//generates new instance of UserController
func NewUserController(userService services.UserService, logger libs.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

//creates new user and associated user address(es)
func (u UserController) CreateNewUser(c *gin.Context) {
	user := entities.User{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.service.WithTrx(trxHandle).CreateUser(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}
