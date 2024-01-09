package mail

import (
	"bez/bez_server/internal/types"
	"log"
	"net/smtp"
)

type IMail struct {
}

func (*IMail) SendMail(to []string, subject string, body string) {
	from := "boazprog@gmail.com"
	password := "qyaw huyx qkhd ayid"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	toStr := "To:"
	for _, t := range to {
		toStr += " " + t + ","
	}
	toStr = toStr[:len(toStr)-1] + "\r\n"

	message := []byte(toStr + "Subject: " + subject + "\r\n" + body + "\r\n")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *IMail) AgoraAgentMail(data types.AgoraAgentResults) {
	subject := "מוצר חדש פורסם"
	body := ""
	body += "שם: " + data.Name + "\r\n"
	body += "קטגוריה: " + data.Category + "\r\n"
	body += "תת קטגוריה: " + data.SubCategory + "\r\n"
	body += "מצב: " + data.Condition + "\r\n"
	body += "פרטים: " + data.Details + "\r\n"
	body += "קישור: " + data.Link + "\r\n"

	m.SendMail([]string{data.Email}, subject, body)
}
