package bot

import (
	"time"

	"github.com/tucnak/telebot"
)

// FakeTelebot is a Telebot mock
type FakeTelebot struct{}

// Listen for messages
func (b *FakeTelebot) Listen(messages chan telebot.Message, timeout time.Duration) {
	messages <- telebot.Message{}
	close(messages)
}

// SendMessage to telegram
func (b *FakeTelebot) SendMessage(chat telebot.Recipient, message string, opts *telebot.SendOptions) error {
	return nil
}
