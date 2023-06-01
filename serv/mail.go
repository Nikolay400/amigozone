package serv

import (
	"amigozone/c"
	"fmt"
	"net/smtp"
)

var Auth smtp.Auth

func AuthInit() {
	Auth = smtp.PlainAuth("", c.Mail, c.Pw, "smtp.gmail.com")
}

func MailSender(to string, subject string, body string) {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + ".\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", Auth, c.Mail, []string{to}, msg)

	if err != nil {
		fmt.Println(err)
		return
	}
}
