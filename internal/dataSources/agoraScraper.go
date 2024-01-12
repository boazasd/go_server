package dataSources

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"log"
	"strings"
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
			return
		}

		name := e.ChildText(".objectName a")
		condition := e.ChildAttr(".objectState", "title")
		date := e.ChildAttr(".regDate", "title")

		c2 := colly.NewCollector()
		c2.OnHTML("#content", func(e *colly.HTMLElement) {
			details := e.ChildText("table.objectDetails td.details")
			area := e.ChildText("table.objectDetails td.leftSection ul li:first-child")
			image := e.ChildAttr("table.objectDetails td.photoIcon a", "href")
			city := e.ChildText("table.objectDetails td.leftSection ul li:nth-child(2)")
			category := e.ChildText("#pinkPageTitle span a:first-child")
			middleCategory := e.ChildText("#pinkPageTitle span a:nth-child(2)")
			subCategory := e.ChildText("#pinkPageTitle span a:nth-child(3)")

			// "Jan 2, 2006 at 3:04pm (MST)"
			parsedDate, err := time.Parse("2/1/2006 15:04", date)
			utcDate := parsedDate.UTC()

			if err != nil {
				log.Println(err.Error())
				return
			}

			agoraData := types.AgoraData{
				Link:           link,
				Name:           name,
				Condition:      condition,
				Date:           utcDate,
				Details:        details,
				Category:       category,
				MiddleCategory: middleCategory,
				SubCategory:    subCategory,
				Image:          image,
				Area:           strings.Replace(area, "אזור: ", "", 1),
				City:           city,
				Processed:      false,
			}

			_, err = am.CreateAgoraData(agoraData)
			if err != nil {
				log.Println(err.Error())
			}
		})
		c2.Visit(websiteUrl + link)
	})

	c1.Visit(websiteUrl + "/toGet.asp?dealType=1")
}
