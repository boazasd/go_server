package main

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/routes"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	models.CreateDatabase()
	models.ConnectDatabse()
	models.CreateFirstUser()
	// runDev()

	defer models.CloseDatabase()
	routes.Init()
}
