package bot

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	photo = "photo"
)

// Process telegram update
func (d *Dialog) Process(update tgbotapi.Update) []tgbotapi.Chattable {
	var result []tgbotapi.Chattable
	result = append(result, d.next(update)...)
	return result
}

// Then takes handler which called when waiter works
func (d *Dialog) Then(f MessageHandler) Processor {
	log.Debug(fmt.Sprintf("Handler saved: %v", f))
	d.next = f
	return d
}

// WaitFor takes waiter to wait in dialog
func (d *Dialog) WaitFor(waiter Waiter) Waiter {
	log.Debug(fmt.Sprintf("Waiter saved: %v", waiter))
	d.waiter = waiter
	return d
}

// Waited checks waiter for wait event
func (d *Dialog) Waited(message *tgbotapi.Message) (bool, string) {
	return false, ""
}
