package bot

import (
	"time"

	"github.com/tucnak/telebot"
)

// Bot implements bot
type Bot interface {
	Sender
	Listener
}

// Database describes bot db operations
type Database interface {
	GetIDByUsername(messenger, username string) (string, error)
	SetUsernameID(messenger, username, id string) error
	RegisterBotIfAlreadyNot(string) bool
	DeregisterBots()
	DeregisterBot(string) error
}

// Listener starts bot
type Listener interface {
	Listen()
}

// Sender sends message
type Sender interface {
	Send(login, message string) error
}

// Telebot interface for telegram bot
type Telebot interface {
	Listen(chan telebot.Message, time.Duration)
	SendMessage(telebot.Recipient, string, *telebot.SendOptions) error
}
