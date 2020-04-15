package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/emersion/go-imap"
	idle "github.com/emersion/go-imap-idle"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"

	"log"
)

type MailClient struct {
	Client *client.Client
	Idle MailClientIdle

}
type MailClientIdle struct {
	IdleErrors chan error
	IdleUpdates chan client.Update
	IdleStop chan struct{}
}

func (ec *MailClient) InitIdle() {
	// Create IDLE client
	idleClient := idle.NewClient(ec.Client)
	// Create a channel to receive mailbox updates
	updates := make(chan client.Update)
	ec.Client.Updates = updates
	// Start idling
	idleErr := make(chan error, 1)
	idleStop:= make(chan struct{}, 1)
	go func() {
		idleErr <- idleClient.IdleWithFallback(nil, 0)
	}()
	ec.Idle.IdleErrors = idleErr
	ec.Idle.IdleUpdates = updates
	ec.Idle.IdleStop = idleStop
}


func (ec *MailClient) ListenForEmails() {
	// Listen for updates
	for {
		select {
		case update := <-ec.Idle.IdleUpdates:
			switch  u := update.(type) {
			case *client.MailboxUpdate:
				//ec.Idle.IdleStop <- struct {}{} // Fetches fine, short read error when trying to fetch again
				//close(ec.Idle.IdleStop) // short read error upon trying to fetch as well as when trying to re-idle
				log.Println("Event MailboxUpdate")
				messages := make(chan *imap.Message, 10)
				seq := new(imap.SeqSet)
				seq.AddNum(u.Mailbox.Messages)
				go func() {
					err := ec.Client.Fetch(seq, []imap.FetchItem{imap.FetchRFC822}, messages)
					if err == nil {
						for email := range messages {
							log.Println("New mail: ")
							for _, body := range email.Body {
								go MailProcessing(ec.EmailBodyParse(body), email)
							}
						}
					} else {
						// Deal with error
						log.Println("Fetch error: ", err)
					}
				}()
			default:
				log.Println("New update: \n", spew.Sdump(update))
			}
		case err := <-ec.Idle.IdleErrors:
			if err != nil {
				log.Println(err)
			}
			log.Println("Not idling anymore")
			// re-idle
			ec.InitIdle()
		}
	}
}
func (ec *MailClient) EmailBodyParse(r io.Reader) []byte {
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			log.Println("got text: ", string(b))
			return b
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			// Unused in my case
			log.Println("got attachment: ", filename)
		}
	}
	return []byte{}
}