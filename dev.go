package main

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"os"
)

func dev() {

	// dataSources.ScrapeAgora()
	// m := mail.IMail{}
	// m.SendMail([]string{"boazasd@gmail.com"}, "מוצר חדש פורסם", "כורסא במצב טוב מאזור  תל אביב ")
	agoraData := []types.AgoraData{
		{Name: "asd"},
		{Name: "lalal"},
	}
	um := models.IUser{}
	res, err := um.GetForAgentMessage(agoraData)

	if err != nil {
		panic(err)
	}

	for _, user := range res {
		println(user.FirstName, user.LastName, user.Email, user.SearchTxt)
	}

}

func runDev() {
	if true {
		dev()
		os.Exit(0)
	}
}
