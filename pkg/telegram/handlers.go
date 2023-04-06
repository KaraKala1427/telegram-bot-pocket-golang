package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"log"
)

const (
	commandStart = "start"

	startReply = "Привет, чтобы сохранить ссылку на своем Pocket аккаунте, для начала тебе неодходимо дать мне на это доступ. Для этого переходи по ссылке: \n%s"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Бот активировался")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "я не знаю эту комманду")
	_, err := b.bot.Send(msg)
	return err
}
