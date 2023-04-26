package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"tg-bot/entity"
	"tg-bot/tg"
	"tg-bot/tg/handlers"
	"tg-bot/tg/users"
)

func main() {

	// Загрузка токена
	var err error
	file, err := os.ReadFile("config.txt")
	if err != nil {
		log.Fatalln(err)
	}
	token := string(file)

	// Инициализация бота и передача ему токена
	bot, err := tg.Init(token)
	if err != nil {
		log.Fatalln(err)
	}

	// Добавление админов
	users.AddUser(entity.NewAdmin("m_m_abdul", "admin"))
	users.AddUser(entity.NewAdmin("User_Anastasia", "admin"))

	// Регистрируем события админа
	handlers.LoadAdminHandlers()

	// Логирование всех действий
	bot.Debug = true

	// Конфигурация канала обновлений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// Запуск прослушивания обновлений от бота
	log.Printf("%s started", bot.Self.UserName)
	tg.Listen(bot, updateConfig)
}
