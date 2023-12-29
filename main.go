package main

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/routes"
)

func main() {
	runDev()
	models.CreateDatabase()
	models.ConnectDatabse()
	defer models.CloseDatabase()
	routes.Init()
}
