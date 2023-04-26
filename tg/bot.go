package tg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/tg/handlers"
)

func Init(token string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(token)
}

func Listen(bot *tgbotapi.BotAPI, config tgbotapi.UpdateConfig) {
	handlers.Init(bot)
	updates := bot.GetUpdatesChan(config)
	for update := range updates {
		for _, handler := range handlers.GetHandlers() {
			//handler.AsyncLoad(&update)
			if handler.Load(&update) {
				break
			}
		}
	}
}
