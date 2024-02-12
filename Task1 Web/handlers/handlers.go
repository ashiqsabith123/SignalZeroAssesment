package handlers

import (
	"fmt"
	"singnalzero-assesment/models"
	"singnalzero-assesment/repository"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userDetails models.User

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"code":    400,
			"message": "data is not in proper format ",
			"error":   "data error",
		})
		return
	}

	if userDetails.Name == "" || userDetails.Age == 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"code":    400,
			"message": "name and age required",
			"error":   "data error",
		})
		return
	}

	err := repository.CreateUser(c, userDetails)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    500,
			"message": "error while inserting data",
			"error":   "server error",
		})
		return
	}

	c.JSON(201, gin.H{
		"code":    201,
		"message": "user created succesfully",
	})

}

func GetAllUsers(c *gin.Context) {

	users, err := repository.GetAllUsers(c)

	fmt.Println(err)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    500,
			"message": "error while fetching data",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "users data fetched succesfully",
		"data":    users,
	})

	return

}

func GetUserByName(c *gin.Context) {

	name := c.Query("name")

	if name == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"code":    400,
			"message": "query name required",
			"error":   "invalid",
		})

		return
	}

	user, err := repository.GetUserByName(c, name)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    500,
			"message": "error while fetching data",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code ":   200,
		"message": "user data fetched succesfully",
		"data":    user,
	})

}
