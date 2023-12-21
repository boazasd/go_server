package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init() {
	router.GET("/v", func(c *gin.Context) {
		c.String(200, "v0.0.2")
	})
	// component := templates.Hello("world")
	// router.GET("/hello", templ.Handler(component))
	usersInit()
	router.Run(":8080")
}
