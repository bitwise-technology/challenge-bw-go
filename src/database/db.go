package database

import (
	"fmt"

	"github.com/challenge-bw-go/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "bw"
const DB_PASSWORD = "bw"
const DB_NAME = "bw_go"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}

func Initialize() {
	db := InitDb()
	db.AutoMigrate(&models.User{})
}
