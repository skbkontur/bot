package bot

import (
	"github.com/tucnak/telebot"
	"gopkg.in/telegram-bot-api.v4"
)

// Dialog handles dialogs between bot and user
type Dialog struct {
	ID        int64
	Room      *Room
	moderator int64
	name      string
	waiter    Waiter
	next      MessageHandler
}

// Dialogs is a map of dialogs
type Dialogs map[int64]*Dialog

// Room handles dialogs between group users and bot
type Room struct {
	ID        int64
	name      string
	moderator Moderator
	dialogs   Dialogs
}

// Rooms is a map of rooms
type Rooms map[int64]*Room

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
