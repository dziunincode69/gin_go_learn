package models

import (
	"gin_go_learn/internal/helper"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"not null"`
}

func (U *User) BeforeCreate(db *gorm.DB) error {
	U.Password = helper.HashPassword(U.Password)
	return nil
}
