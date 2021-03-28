package controllers

import (
	"github.com/challenge-bw-go/src/database"
	"github.com/challenge-bw-go/src/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}