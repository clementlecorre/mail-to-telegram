package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var (
	b *tb.Bot
	userID = &tb.Chat{ID: config.TelegramUserID}
)

type MessageFmt struct {
	Subject string
	Link string

}
func init() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		Token:  config.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("telegram: ", err)
		return
	}
}


func MessageFormatting(msgFmt MessageFmt) string {
	return fmt.Sprintf("%s\n" +
		"%s",msgFmt.Subject, msgFmt.Link)
}
