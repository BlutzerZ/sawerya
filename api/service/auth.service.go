package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"blutzerz/sawerya/helpers"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		db: configs.GetDB(),
	}
}

func (as *AuthService) Login(req *dto.LoginUserRequest) (string, error) {
	var user models.User

	err := as.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	acessToken, err := helpers.GenerateAccessToken(user.ID, user.Username)

	return acessToken, err
}
