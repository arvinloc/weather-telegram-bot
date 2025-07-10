package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"weather-tg-bot/internal/api"
	"weather-tg-bot/internal/config"
	"weather-tg-bot/internal/telegram/handlers"
)

func main() {
	cfg := config.LoadConfig()

	client := api.NewWeatherClient(cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}
	handler := handlers.NewTelegramHandler(bot, client)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			handler.HandleCommand(update.Message)
			continue
		}
		handler.HandleText(update.Message)
	}

}
