package dataSources

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"log"
	"time"

	"github.com/gocolly/colly"
)

const websiteUrl = "https://www.agora.co.il"

func ScrapeAgora() {
	c1 := colly.NewCollector()

	c1.OnHTML("tbody.objectGroup", func(e *colly.HTMLElement) {
		am := models.IAgoraModel{}
		link := e.ChildAttr(".newWindow a[href]", "href")
		isExist, err := am.GetAgoraDataByLink(link)
		if err != nil {
			log.Println(err.Error())
		}

		if isExist.Link != "" {
			log.Println("already exist")
			return
		}
		log.Println(link)
		name := e.ChildText(".objectName a")
		state := e.ChildAttr(".objectState", "title")
		date := e.ChildAttr(".regDate", "title")

		c2 := colly.NewCollector()
		c2.OnHTML("table.objectDetails", func(e *colly.HTMLElement) {
			details := e.ChildText("td.details")
			area := e.ChildText("td.leftSection ul li:first-child")
			city := e.ChildText("td.leftSection ul li:nth-child(2)")

			// "Jan 2, 2006 at 3:04pm (MST)"
			parsedDate, err := time.Parse("2/1/2006 15:04", date)

			if err != nil {
				log.Println(err.Error())
				return
			}

			agoraData := types.AgoraData{
				Link:    link,
				Name:    name,
				State:   state,
				Date:    parsedDate,
				Details: details,
				Area:    area,
				City:    city,
			}
			id, err := am.CreateAgoraData(agoraData)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(id)
		})
		c2.Visit(websiteUrl + link + "?toGet=1")
	})

	c1.Visit(websiteUrl + "/toGet.asp?dealType=1")
	log.Println("done")
}
