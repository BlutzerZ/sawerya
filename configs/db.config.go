package configs

import (
	"blutzerz/sawerya/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB() error {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/sawerya?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	dbClient.AutoMigrate(&models.User{}, &models.Alert{}, models.AlertDesign{}, &models.TransactionType{}, &models.Transaction{})

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}
