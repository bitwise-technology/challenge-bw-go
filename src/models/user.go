package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID int
	Username string
	Name string
	LastName string
	ProfileImageUrl string
	Bio string
	Email string 
	Gender string
}