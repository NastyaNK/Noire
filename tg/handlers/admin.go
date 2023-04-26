package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"tg-bot/entity"
	"tg-bot/tg/users"
)

func LoadAdminHandlers() {
	adminMainKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать сущность", "admin_create_entity"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать карту", "admin_create_card"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сёгуны", "admin_list_"+entity.SHOGUN),
			tgbotapi.NewInlineKeyboardButtonData("Дайме", "admin_list_"+entity.DAIMYO),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Самураи", "admin_list_"+entity.SAMURAI),
			tgbotapi.NewInlineKeyboardButtonData("Инкассаторы", "admin_list_"+entity.COLLECTOR),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Карты", "admin_list_card"),
		),
	)
	adminCreateKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сёгун", "admin_create_entity_"+entity.SHOGUN),
			tgbotapi.NewInlineKeyboardButtonData("Дайме", "admin_create_entity_"+entity.DAIMYO),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Самурай", "admin_create_entity_"+entity.SAMURAI),
			tgbotapi.NewInlineKeyboardButtonData("Инкассатор", "admin_create_entity_"+entity.COLLECTOR),
		),
	)
	AddCommandHandler(
		"functions",
		func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
			user.SetMode("")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Функции администратора:")
			msg.ReplyMarkup = adminMainKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return false
		},
	)
	// Создание сущности
	AddQueryDataHandler(
		"admin_create_entity",
		func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
			user.SetMode("")
			msg := tgbotapi.NewEditMessageTextAndMarkup(
				update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID,
				"Создание сущностей",
				adminCreateKeyboard,
			)
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			return false
		},
	)
	AddContainsQueryDataHandler("admin_create_entity_", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		user.SetMode("")
		if _, ok := user.GetEntity().(*entity.Admin); ok {
			query := strings.Split(update.CallbackData(), "_")
			entityType := query[len(query)-1]
			user.SetMode("create_entity_username_" + entityType)
			_, _ = bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите username пользователя"))
			return true
		}
		_, _ = bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вы не можете добавить сущность"))
		return true
	})
	AddMessageContainsModeHandler(
		"create_entity_username_",
		func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
			user.SetData("entity_username", update.Message.Text)
			user.SetMode(strings.Replace(user.Mode(), "username", "nickname", 1))
			_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Введите nickname пользователя"))
			return true
		},
	)
	
	AddMessageContainsModeHandler(
		"create_entity_nickname_",
		func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
			user.SetData("entity_nickname", update.Message.Text)
			if admin, ok := user.GetEntity().(*entity.Admin); ok {
				query := strings.Split(user.Mode(), "_")
				entityType := query[len(query)-1]
				username, ok := user.GetData("entity_username")
				if !ok {
					_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не ввели username"))
					user.SetMode("create_entity_username_" + entityType)
					return true
				}
				nickname, ok := user.GetData("entity_nickname")
				if !ok {
					_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не ввели nickname"))
					user.SetMode("create_entity_nickname_" + entityType)
					return true
				}
				newEntity, err := admin.NewEntity(entityType, username, nickname)
				if err != nil {
					_, _ = bot.Send(
						tgbotapi.NewMessage(
							update.Message.Chat.ID,
							err.Error(),
						),
					)
					return false
				}
				users.AddUser(newEntity)
				_, _ = bot.Send(
					tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Новый пользователь добавлен",
					),
				)
				return true
			}
			return true
		},
	)
	// Создание выбранной сущности

	AddQueryDataHandler("admin_create_card", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Создание карты",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
	AddQueryDataHandler("admin_list_shogun", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Список сёгунов",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
	AddQueryDataHandler("admin_list_daimyo", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Список дайме",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
	AddQueryDataHandler("admin_list_samurai", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Список самураев",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
	AddQueryDataHandler("admin_list_collector", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Список инкассаторов",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
	AddQueryDataHandler("admin_list_card", func(bot *tgbotapi.BotAPI, handlerId int, update *tgbotapi.Update, user *users.User) bool {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			"Список карт",
			adminMainKeyboard,
		)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		return false
	})
}
