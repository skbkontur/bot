package bot

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

const messenger = "telegram"

// Start api loop
func (b *telegram) Loop(messages chan telebot.Message, timeout time.Duration) {
	b.bot.Listen(messages, timeout)

	for message := range messages {
		if err := b.handleMessage(message); err != nil {
			log.WithError(err).Error("Error sending message")
		}
	}
}

func (b *telegram) Talk(username, message string) error {
	uid, err := b.db.GetIDByUsername(messenger, username)
	if err != nil {
		return err
	}
	var options *telebot.SendOptions
	return b.bot.SendMessage(recipient{uid}, message, options)
}

func (b *telegram) handleMessage(message telebot.Message) error {
	var err error
	var options *telebot.SendOptions
	id := strconv.FormatInt(message.Chat.ID, 10)
	title := message.Chat.Title
	userTitle := strings.Trim(fmt.Sprintf("%s %s", message.Sender.FirstName, message.Sender.LastName), " ")
	username := message.Chat.Username
	chatType := message.Chat.Type
	switch {
	case chatType == "private" && message.Text == "/start":
		if username == "" {
			b.bot.SendMessage(message.Chat, "Username is empty. Please add username in Telegram.", options)
		} else {
			err = b.db.SetUsernameID(messenger, "@"+username, id)
			if err != nil {
				return err
			}
			b.bot.SendMessage(message.Chat, fmt.Sprintf("Okay, %s, your id is %s", userTitle, id), nil)
		}
	case chatType == "supergroup" || chatType == "group":
		fmt.Println(chatType, title)
		err = b.db.SetUsernameID(messenger, title, id)
		if err != nil {
			return err
		}
		b.bot.SendMessage(message.Chat, fmt.Sprintf("Hi, all!\nI will send alerts in this group (%s).", title), nil)
	default:
		b.bot.SendMessage(message.Chat, "I don't understand you :(", nil)
	}
	return err
}
