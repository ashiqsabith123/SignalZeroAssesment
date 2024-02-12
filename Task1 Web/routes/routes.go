package routes

import (
	"singnalzero-assesment/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/create", handlers.CreateUser)
		user.GET("/get/all", handlers.GetAllUsers)
		user.GET("/get/user", handlers.GetUserByName)
	}

}
