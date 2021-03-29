package controllers

import (
	"net/http"

	"github.com/challenge-bw-go/src/models"
	"github.com/challenge-bw-go/src/repositories"
	"github.com/challenge-bw-go/src/validators"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if validators.UserCreateValidator(input) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user := models.User{
		Username: input.Username, Name: input.Name, LastName: input.LastName, ProfileImageUrl: input.ProfileImageUrl,
		Bio: input.Bio, Email: input.Email, Gender: input.Gender,
	}

	if err := repositories.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	users := repositories.GetUsers()

	c.JSON(http.StatusCreated, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.VerifyIfUserExists(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var user models.User

	if err := repositories.GetUser(&user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
