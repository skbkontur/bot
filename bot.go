package bot

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	log   Logger
	debug = true
)

func (b *telegram) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.api.GetUpdatesChan(u)

	var messages []tgbotapi.Chattable
	for update := range updates {
		messages = b.processor.Process(update)
		for _, msg := range messages {
			b.api.Send(msg)
		}
	}
}

// StartTelegramBot init telegram api bot
func StartTelegramBot(token string, processor Processor) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error(err)
	}
	api.Debug = debug
	log.Debug(fmt.Sprintf("Authorized on account %s", api.Self.UserName))

	//tg := telegram{api: api}
	//go tg.api.Listen()
}
