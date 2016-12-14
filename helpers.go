package bot

import (
	"time"

	"github.com/tucnak/telebot"
)

// StartTelebot creates an api and start telebot
func StartTelebot(key string, db Database) (Bot, error) {
	messages := make(chan telebot.Message)

	api := &telegram{
		key: key,
		db:  db,
	}
	var err error
	api.bot, err = telebot.NewBot(key)
	if err == nil && db.RegisterBotIfAlreadyNot(key) {
		go api.Loop(messages, 1*time.Second)
	}
	return api, err
}
