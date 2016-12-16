package bot

import (
	"time"

	"github.com/tucnak/telebot"
	"gopkg.in/telegram-bot-api.v4"
)

// Bot implements bot
type Bot interface {
	Talker
}

// API implements bot go-telegram-bot-api interface
type API interface {
	GetUpdatesChan(u tgbotapi.UpdateConfig) (<-chan tgbotapi.Update, error)
	Listener
	Sender
}

// Database describes bot db operations
type Database interface {
	GetIDByUsername(messenger, username string) (string, error)
	SetUsernameID(messenger, username, id string) error
	RegisterBotIfAlreadyNot(string) bool
	DeregisterBots()
	DeregisterBot(string) error
}

// Listener starts bot long-polling
type Listener interface {
	Listen(chan Message, time.Duration)
}

// Logger logs messages
type Logger interface {
	WithError(err error) Logger
	WithFields(str interface{}) Logger
	Debug(message ...interface{})
	Error(message ...interface{})
	Fatal(message ...interface{})
	Info(message ...interface{})
}

// Message is a interface for cross-package messages
type Message interface {
	Chattable() tgbotapi.Chattable
}

// Messages implements chan of message
type Messages chan Message

// MessageHandler is a func for bot api message handling
type MessageHandler func(update tgbotapi.Update) []tgbotapi.Chattable

// Moderator runs dialog on moderator side
type Moderator interface {
	//Moderate(message *tgbotapi.Message, dialog *Dialog) []tgbotapi.Chattable
	ModeratorID() int64
}

// Processor processes bot updates
type Processor interface {
	Process(update tgbotapi.Update) []tgbotapi.Chattable
}

// Starter starts process, usually as goroutine
type Starter interface {
	Start()
}

// Sender sends a chattable
type Sender interface {
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
}

// Talker talks to username
type Talker interface {
	Talk(username, message string) error
}

// Telebot interface for telegram bot
type Telebot interface {
	SendMessage(telebot.Recipient, string, *telebot.SendOptions) error
}

// Waiter waits for
type Waiter interface {
	Waited(message *tgbotapi.Message) (bool, string)
}
