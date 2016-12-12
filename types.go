package bot

import "github.com/tucnak/telebot"

type telegram struct {
	key      string
	telebot  Telebot
	messages chan telebot.Message
	db       Database
}

type recipient struct {
	uid string
}
