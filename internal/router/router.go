package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init() {
	fmt.Println("Hello World")
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Run(":8080")
}