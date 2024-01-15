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
	schedule()

	defer models.CloseDatabase()

	routes.Init()
}

func schedule() {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				dataSources.ScrapeAgora()
				am := models.IAgoraModel{}
				data, err := am.GetForAgentMessage()
				log.Println("items for agents", len(data))
				if err != nil {
					log.Println(err)
				}
				m := mail.IMail{}
				for _, d := range data {
					log.Println(d.Email, d.AgentId)
					go m.AgoraAgentMail(d)
				}
				am.UpdateProcessed()
			case <-quit:
				println("stop tick")
				ticker.Stop()

				return
			}
		}
	}()
}
