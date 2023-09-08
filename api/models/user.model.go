package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"type:integer; primaryKey"`
	Username  string `gorm:"type:string; size:20; not null; unique"`
	Email     string `gorm:"type:string; size:30; not null; unique"`
	Password  string `gorm:"type:string; size:72; not null"`
	CreatedAt uint   `gorm:"type:integer; not null"`
	UpdatedAt uint   `gorm:"type:integer; not null"`
	Alert     *Alert
}

// Before
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Password", hash)

	// generate time
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		// hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		tx.Statement.SetColumn("Password", hash)
	}
	// generate time
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

// After
func (u *User) AfterCreate(tx *gorm.DB) error {
	alert := Alert{
		EnableGif:       0,
		MinAmountNotify: 5000,
		MinAmountGIF:    5000,
		Sound:           "default",
		UserID:          u.ID,
	}
	AlertDesign := AlertDesign{
		BackgroundColor: "#199999",
		HighlightColor:  "#000000",
		TextColor:       "#000000",
		TextTemplate:    "baru saja memberikan",
		Border:          0,
		TextTickness:    100,
		Duration:        5,
		Font:            "arial",
	}

	alert.AlertDesign = &AlertDesign
	err := tx.Model(Alert{}).Create(&alert).Error

	if err != nil {
		return err
	}

	return nil
}
