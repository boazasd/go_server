package main

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/routes"
)

func main() {
	models.CreateDatabase()
	models.ConnectDatabse()
	routes.Init()
}