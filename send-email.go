package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
)

var (
	msg      = flag.String("msg", "", "The message to send.")
	from     = flag.String("from", "", "Email to log in with.")
	password = flag.String("password", "", "Password to sign in with.")
	to       = flag.String("to", "", "Email to send to.")
)

func main() {
	flag.Parse()
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", *from, *password, smtpHost)
	headers := fmt.Sprintf("From: raspberrypi <%s>\nSubject: Message from raspberry pi\nTo: %s\n\n",
		*from, *to)
	if *msg == "" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		s := string(b)
		msg = &s
	}
	message := append([]byte(headers), []byte(*msg)...)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, *from, []string{*to}, message)
	if err != nil {
		log.Fatal(err)
	}
}
