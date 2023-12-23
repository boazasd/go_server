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
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.String(200, "")
	})
	component := templates.Login()
	router.GET("/", gin.WrapH(templ.Handler(component)))

	usersInit()
	router.Run(":8080")
}