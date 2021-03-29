package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              int    `gorm:"primaryKey"`
	Username        string `json:"username" gorm:"unique;not null" validate:"required,alphanum"`
	Name            string `json:"name" gorm:"not null" validate:"required,alpha"`
	LastName        string `json:"lastname" gorm:"default:null"`
	ProfileImageUrl string `json:"profileImageUrl" gorm:"default:null"`
	Bio             string `json:"bio" gorm:"default:null"`
	Email           string `json:"email" gorm:"unique;not null" validate:"email"`
	Gender          string `json:"gender" gorm:"default:null"`
}
