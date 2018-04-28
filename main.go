package main

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := api.NewBotAPI("571704538:AAEAGONOB5-tWBGz_uqDrXTuDUYBMKfW5Lk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	cfg := api.NewUpdate(0)
	cfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(cfg)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}