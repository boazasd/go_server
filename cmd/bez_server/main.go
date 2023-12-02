package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("Hello World")
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Run(":8080")
}
