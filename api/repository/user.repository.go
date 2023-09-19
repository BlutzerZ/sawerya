package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: configs.GetDB(),
	}
}

func (ur *UserRepository) Create(user models.User) error {
	tx := ur.DB.Begin()

	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (ur *UserRepository) FindByID(ID uint) (models.User, error) {
	var user models.User

	err := ur.DB.Where("id = ?", ID).First(&user).Error

	return user, err
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := ur.DB.Find(&users).Error

	return users, err
}

func (ur *UserRepository) Update(user *models.User) error {
	tx := ur.DB.Begin()

	err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ur *UserRepository) Delete(ID uint) error {
	tx := ur.DB.Begin()

	err := tx.Model(&models.User{}).Where("id = ?", ID).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
