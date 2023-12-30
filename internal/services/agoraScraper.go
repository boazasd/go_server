package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"time"

	"github.com/gocolly/colly"
)

const websiteUrl = "https://www.agora.co.il"

func ScrapeAgora() {
	// arr := make(map[string]types.AgoraData)

	c1 := colly.NewCollector()

	c1.OnHTML("tbody.objectGroup", func(e *colly.HTMLElement) {

		link := e.ChildAttr(".newWindow a[href]", "href")
		isExist, err := models.GetAgoraDataByLink(link)
		if err != nil {
			println(err.Error())
		}

		if isExist.Link != "" {
			return
		}
		println(link)
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
				println(err.Error())
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
			id, err := models.CreateAgoraData(agoraData)
			if err != nil {
				println(err.Error())
			}
			println(id)
		})
		c2.Visit(websiteUrl + link + "?toGet=1")

		// println(id)

		// println(name)
		// println(state)
		// println(city)
		// println(details)
		// println(date)
		// println(area)
	})

	c1.Visit(websiteUrl + "/toGet.asp?dealType=1")
	println("done")
}
