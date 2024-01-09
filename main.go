package main

import (
	"bez/bez_server/internal/dataSources"
	"bez/bez_server/internal/mail"
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/routes"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	models.CreateDatabase()
	models.ConnectDatabse()
	models.CreateFirstUser()
	runDev()

	defer models.CloseDatabase()
	routes.Init()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				println("tick")
				dataSources.ScrapeAgora()
				am := models.IAgoraModel{}
				am.UpdateProcessed()
				data, err := am.GetForAgentMessage()
				if err != nil {
					log.Println(err)
					return
				}
				m := mail.IMail{}
				for _, d := range data {
					m.AgoraAgentMail(d)
				}
				return
			case <-quit:
				println("stop tick")
				ticker.Stop()

				return
			}
		}
	}()
}
