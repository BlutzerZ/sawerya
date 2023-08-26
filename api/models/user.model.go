package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"type:integer; primaryKey"`
	Username  string `gorm:"type:string; size:20; not null; unique"`
	Email     string `gorm:"type:string; size:30; not null; unique"`
	Password  string `gorm:"type:string; size:72; not null"`
	CreatedAt int    `gorm:"type:integer; not null"`
	UpdatedAt int    `gorm:"type:integer; not null"`
}

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
