package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/karakala1427/telegram-bot-pocket-golang/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6108718086:AAF7fowK_buhRHof7VsCPobZw18I8q8OR9A")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("106726-32d8215d8a78053afc25080")

	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
