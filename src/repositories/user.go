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

func GetUser(user *models.User, id string) error {
	if err := database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, input *models.User) error{
	if err := database.Db.Model(&user).Updates(input).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User, id string) error {
	if err := database.Db.Delete(user, id).Error; err != nil {
		return err
	}
	return nil
}