package routes

import (
	"bez/bez_server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init() {
	router.GET("/v", func(c *gin.Context) {
		c.String(200, "v0.0.2")
	})
	component := templates.Hello("Shirit")
	router.GET("/hello", gin.WrapH(templ.Handler(component)))

	usersComponent := templates.Users([]string{"one", "two", "three"})
	router.GET("/users", gin.WrapH(templ.Handler(usersComponent)))

	usersInit()
	router.Run(":8080")
}
