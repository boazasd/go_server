package main

import (
	"bez/bez_server/internal/dataSources"
	"os"
)

func dev() {

	dataSources.ScrapeAgora()

}

func runDev() {
	if true {
		dev()
		os.Exit(0)
	}
}
