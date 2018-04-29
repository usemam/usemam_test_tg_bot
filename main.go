package main

import (
	"log"
	"net/http"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	config "github.com/usemam/usemam_test_tg_bot/configuration"
)

func processUpdate(update api.Update, bot *api.BotAPI) error {
	if update.Message == nil {
		return nil
	}

	log.Println("New message - '%s' %s", update.Message.From.FirstName, update.Message.Text)
	message := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
	_, err := bot.Send(message)
	return err
}

func main() {
	cfg := config.New()

	bot, err := api.NewBotAPI(cfg.BotToken)
	fail(err)

	_, err = bot.SetWebhook(api.NewWebhook(cfg.URL + cfg.BotToken))
	fail(err)

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+cfg.Port, nil)

	for update := range updates {
		err = processUpdate(update, bot)
		if err != nil {
			log.Println(err)
		}
	}
}
