package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	gin "github.com/gin-gonic/gin"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot      *api.BotAPI
	botToken string
	baseURL  string
)

func initBot() {
	var err error

	bot, err = api.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
		return
	}

	_, err = bot.SetWebhook(api.NewWebhook(baseURL + botToken))
}

func webHookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update api.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
		return
	}

	botToken = os.Getenv("TOKEN")
	if botToken == "" {
		log.Fatal("$TOKEN must be set")
		return
	}

	baseURL = os.Getenv("URL")
	if baseURL == "" {
		log.Fatal("$URL must be set properly")
	}

	router := gin.New()
	router.Use(gin.Logger())

	initBot()
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	router.POST("/"+bot.Token, webHookHandler)
	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
