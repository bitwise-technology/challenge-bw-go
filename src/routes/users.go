package routes

import (
	"github.com/challenge-bw-go/src/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(route *gin.RouterGroup) {
	group := route.Group("/users")
	group.POST("/", controllers.CreateUser)
	group.GET("/", controllers.GetUsers)
	group.GET("/:id", controllers.GetUser)
	group.PATCH("/:id", controllers.UpdateUser)
	group.DELETE("/:id", controllers.DeleteUser)


}
