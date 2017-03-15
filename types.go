package bot

import (
	"github.com/tucnak/telebot"
	"gopkg.in/telegram-bot-api.v4"
)

type recipient struct {
	uid string
}

type telegram struct {
	api       *tgbotapi.BotAPI
	bot       *telebot.Bot
	db        Database
	key       string
	messages  chan Message
	processor Processor
}
