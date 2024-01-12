package mail

import (
	"bez/bez_server/internal/types"
	"log"
	"net/smtp"
)

type IMail struct {
}

func (*IMail) SendMail(to []string, subject string, body string, html bool) {
	from := "boazprog@gmail.com"
	password := "qyaw huyx qkhd ayid"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	toStr := "To:"
	for _, t := range to {
		toStr += " " + t + ","
	}
	toStr = toStr[:len(toStr)-1] + "\r\n"
	mime := "\n"
	if html {
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	}
	message := []byte(toStr + "Subject: " + subject + "\r\n" + mime + body + "\r\n")

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
	body := "<div dir=rtl >"
	body += "שם: " + data.Name + "<br/>"
	body += "קטגוריה: " + data.Category + "<br/>"
	body += "תת קטגוריה: " + data.SubCategory + "<br/>"
	body += "מצב: " + data.Condition + "<br/>"
	body += "פרטים: " + data.Details + "<br/>"
	body += "קישור: https://www.agora.co.il" + data.Link + "<br/>"
	body += "</div>"

	m.SendMail([]string{data.Email}, subject, body, true)
}
