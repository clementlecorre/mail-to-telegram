package main

import (
	"github.com/emersion/go-imap"
	"log"
)

func MailProcessing(msg []byte, mail *imap.Message){
	msgFmt := MessageFmt{
		Subject: "",
		Link:    "N/A",
	}
	if mail.Envelope != nil {
		msgFmt.Subject = mail.Envelope.Subject
	}
	msgFmt.Link = MailBodyProcessing(string(msg))
	_, err := b.Send(userID, MessageFormatting(msgFmt))
	if err != nil {
		log.Fatal("telegram: ", err)
		return
	}
}

func MailBodyProcessing(msg string) string {
	// You can customize message processing here
	return "ok"
}