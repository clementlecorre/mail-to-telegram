package main

type Config struct {
	// EmailServer: Example mail.domain.com:993
	EmailServer    string
	EmailLogin     string
	EmailPassword  string
	TelegramUserID int64
	TelegramToken  string
	Verbose        bool
}
