package repositories

import (
	"github.com/challenge-bw-go/src/database"
	"github.com/challenge-bw-go/src/models"
)

func CreateUser(user *models.User) error {
	if err := database.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() []models.User {
	users := []models.User{}
	database.Db.Find(&users)
	return users
}
