package controllers

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/service"
	"blutzerz/sawerya/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		service: service.NewAuthService(),
	}
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	req := new(dto.LoginUserRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "invalid request")
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

	accessToken, err := ac.service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "login failed",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   accessToken,
	})

}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	req := new(dto.RefreshTokenRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "invalid request")
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

	newAccessToken, err := helpers.UpdateToken(req.OldToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err,
			"message": "failed to refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newToken": newAccessToken,
		"message":  "success generate new token",
	})
}
