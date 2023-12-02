package router

import "github.com/gin-gonic/gin"

type role = string

type user struct {
	ID int
	FirstName string
	LastName string
	Email string
	Password string
	Roles []string
}

func init() {
	router.GET ("/users/:id", func(c *gin.Context) {
		user := user{
			ID: 1,
			FirstName: "John",
			LastName: "Doe",
			Email: "jd@bez.com",
		}
		c.JSON(200, user)
	})
}
