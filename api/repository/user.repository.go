package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: configs.GetDB(),
	}
}

func (ur *UserRepository) Create(user models.User) error {
	tx := ur.db.Begin()

	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (ur *UserRepository) FindByID(ID int) error {
	var user models.User

	err := ur.db.Where("id = ?", ID).First(&user).Error

	return err
}

func (ur *UserRepository) FindAll() error {
	var users []models.User

	err := ur.db.Find(&users).Error

	return err
}

func (ur *UserRepository) Update(ID int, user models.User) error {
	tx := ur.db.Begin()

	err := tx.Model(models.User{}).Where("id = ?", ID).Updates(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ur *UserRepository) Delete(ID int) error {
	tx := ur.db.Begin()

	err := tx.Delete(&models.User{}, ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
