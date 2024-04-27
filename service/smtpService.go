package service

import (
	"fmt"
	"net/smtp"
)

func SendMail(from string, to []string, message []byte) {
	host := "app.debugmail.io"
	port := 25
	login := "43a09af1-77d2-40b4-a61d-40fc557c5a02"
	pass := "c81dbb5b-67d2-4ae2-a0ec-5af642b98410"

	auth := smtp.PlainAuth("", login, pass, host)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, login, to, message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Email Sent Successfully")
}
