package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	DB := configs.GetDB()
	return &UserRepository{
		DB: DB,
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

func (ur *UserRepository) FindByID(ID int) (models.User, error) {
	var user models.User

	err := ur.DB.Where("id = ?", ID).First(&user).Error

	return user, err
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := ur.DB.Find(&users).Error

	return users, err
}

func (ur *UserRepository) Update(ID int, field string, value string) error {
	tx := ur.DB.Begin()

	err := tx.Model(&models.User{}).Where("id = ?", ID).Update(field, value).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ur *UserRepository) Delete(ID int) error {
	tx := ur.DB.Begin()

	err := tx.Delete(&models.User{}, ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
