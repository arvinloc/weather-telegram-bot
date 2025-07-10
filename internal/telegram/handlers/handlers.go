package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"weather-tg-bot/internal/models"
)

type TelegramHandler struct {
	bot               *tgbotapi.BotAPI
	weatherProvider   models.WeatherProvider
	awaitingCityInput map[int64]bool
}

func NewTelegramHandler(bot *tgbotapi.BotAPI, weatherProvider models.WeatherProvider) *TelegramHandler {
	return &TelegramHandler{
		bot:               bot,
		weatherProvider:   weatherProvider,
		awaitingCityInput: make(map[int64]bool),
	}
}

func (h *TelegramHandler) HandleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "start":
		msg.Text = "Привет👋, я - твой погодный бот!\n  Используй команду /help, чтобы узнать подробнее😃"
	case "help":
		msg.Text = "⚙️Мои команды:\n/start - начало работы со мной\n/weather - получить погоду желаемого города!🌐"
	case "weather":
		h.awaitingCityInput[message.Chat.ID] = true
		msg.Text = "Введите название города:"
	default:
		msg.Text = "неизвестная команда"
	}
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}

}

func (h *TelegramHandler) HandleText(message *tgbotapi.Message) {
	if h.awaitingCityInput[message.Chat.ID] {
		delete(h.awaitingCityInput, message.Chat.ID)
		data, err := h.weatherProvider.GetWeatherByCity(message.Text)
		if err != nil {
			log.Println(err)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Не удалось получить погоду! Попробуйте другой город")
			h.bot.Send(msg)
			return
		}

		res := fmt.Sprintf("Текущая температура: %.1f\nЧувствуется как: %.1f\nОписание: %s\n",
			data.Main.Temp,
			data.Main.FeelsLike,
			data.Weather[0].Description)
		msg := tgbotapi.NewMessage(message.Chat.ID, res)
		h.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Используйте команду /help")
	h.bot.Send(msg)
}
