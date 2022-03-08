package controllers

import (
	"crypto/hmac"
	"go-todoApp/entities"
	customUtils "go-todoApp/utils"
	userValidator "go-todoApp/validators/userValidator"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "go-todoApp/db"
)

func CreateUser(c *gin.Context) {
	var user userValidator.CreateUser

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errSlice := userValidator.ValidateUserCreation(&user); len(errSlice) == 0 {
		db := db.GetDB()

		var userAddress []entities.Address
		for _, address := range user.Addresses {
			userAddress = append(userAddress, entities.Address{
				Prefecture: address.Prefecture,
				City:       address.City,
				District:   address.District,
				Street:     address.Street,
				Additional: address.Additional,
			})
		}

		userData := entities.User{
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			Addresses: userAddress,
			Role:      entities.Role(user.Role),
		}
		result := db.Create(&userData)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "BadUserInfo",
				"error":   result.Error,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":  "UserAccepted",
				"userData": &userData,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadUserInfo",
			"error":   errSlice,
		})
	}
}

func userLogin(c *gin.Context) {
	session := sessions.Default(c)
	var loginForm userValidator.LoginForm

	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errSlice := userValidator.ValidateUserLogin(&loginForm); len(errSlice) == 0 {
		db := db.GetDB()
		var userData entities.User
		db.Where("username = ?", loginForm.Username).First(&userData)
		if userData.ID == uuid.Nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect Account Info",
			})
		}
		encryptedPassword := customUtils.GenerateHMAC(loginForm.Password, userData.ID.String())
		if hmac.Equal([]byte(userData.Password), []byte(encryptedPassword)) {

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect Account Info",
			})
		}
	} else {

	}
}
