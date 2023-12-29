package main

import (
	"bez/bez_server/internal/services"
	"os"
)

func dev() {

	services.ScrapeAgora()

}

func runDev() {
	if true {
		dev()
		os.Exit(0)
	}
}
