package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              int    `gorm:"primaryKey"`
	Username        string `json:"username" gorm:"unique;not null" binding:"required" validate:"required,alphanum"`
	Name            string `json:"name" gorm:"not null" binding:"required" validate:"required,alpha"`
	LastName        string `json:"lastname" gorm:"default:null"`
	ProfileImageUrl string `json:"profileImageUrl" gorm:"default:null"`
	Bio             string `json:"bio" gorm:"default:null"`
	Email           string `json:"email" gorm:"unique;not null" binding:"required" validate:"email"`
	Gender          string `json:"gender" gorm:"default:null"`
}
