package routes

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/services"
	"bez/bez_server/templates"
	"strconv"

	"github.com/gin-gonic/gin"
)

func usersInit() {
	router.GET("/users/getOne/:id", getUser)
	router.GET("/users/get", getUsers)
	router.POST("/users/create", createUser)
	router.POST("/users/update", updateUser)
	router.DELETE("/users/delete/:id", deleteUser)
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

func getUsers(c *gin.Context){
	users, err := services.GetUsers()
	if err != nil {
		errorComponent := templates.Error(err.Error())
		errorComponent.Render(c.Request.Context(),c.Writer)
	}
	usersComponent := templates.Users(users)
	usersComponent.Render(c.Request.Context(),   
	c.Writer)
}

func createUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	println(user.FirstName, user.LastName, user.Email, user.Password)
	userId, err := services.CreateUser(user)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, userId)
	}
}

func updateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "update user"})
}

func deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "delete user"})
}
