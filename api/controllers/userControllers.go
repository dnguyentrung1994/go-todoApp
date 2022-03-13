package controllers

import (
	"go-todoApp/constants"
	"go-todoApp/entities"
	"go-todoApp/libs"
	"go-todoApp/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	userValidator "go-todoApp/validators/userValidator"
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
	//validator for creating new user
	user := userValidator.CreateUser{}
	// user := entities.User{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if validationErrors := userValidator.ValidateUserCreation(user); len(validationErrors) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Invalid User Data",
			"error lists": validationErrors,
		})
	}
	validatedUserData := entities.User{
		Username: user.Username,
	}
	if err := u.service.WithTrx(trxHandle).CreateUser(validatedUserData); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}
