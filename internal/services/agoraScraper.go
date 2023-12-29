package services

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeAgora() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tbody[class='objectGroup']", func(e *colly.HTMLElement) {
		println(e.ChildText(".newWindow"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.agora.co.il/toGet.asp?dealType=1")
}
