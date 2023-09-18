package controllers

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

func (uc *UserController) CreateNewUser(c *gin.Context) {
	req := new(dto.RegisterUserRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// service
	err = uc.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) UpdateUsername(c *gin.Context) {
	req := new(dto.UpdateUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// service
	err = uc.service.UpdateUsername(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update username",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	req := new(dto.UpdatePasswordRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// service
	err = uc.service.UpdatePassword(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update password",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
