package models

import (
	"gorm.io/gorm"
	//"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	//gorm.Model
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	IsAdmin   bool           `json:"isAdmin" gorm:"default:false"`
	Approved  bool           `gorm:"default:false"`
	Role      string         `gorm:"not null"`
}