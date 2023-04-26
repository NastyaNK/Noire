package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"sync"
	"tg-bot/tg/users"
)

var bot *tgbotapi.BotAPI

func Init(api *tgbotapi.BotAPI) {
	bot = api
}

type HandlerIDGenerator struct {
	mutex *sync.Mutex
	next  int
}

func (g *HandlerIDGenerator) GetNext() int {
	var i int
	g.mutex.Lock()
	g.next++
	i = g.next
	g.mutex.Unlock()
	return i
}

var GeneratorId = &HandlerIDGenerator{&sync.Mutex{}, 0}

type Condition func(*tgbotapi.Update) bool
type Runnable func(*tgbotapi.BotAPI, int, *tgbotapi.Update, *users.User) bool

type Handler struct {
	id        int
	Condition Condition
	Runnable  Runnable
}

func (h *Handler) Load(update *tgbotapi.Update) bool {
	println("HandlerId: before condition", h.id)
	if h.Condition(update) {
		println("HandlerId: after condition", h.id)
		var chatID int64
		var err error
		var user *users.User
		if update.Message != nil {
			chatID = update.Message.Chat.ID
			user, err = users.GetUserByUsername(update.Message.From.UserName)
		}
		if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
			user, err = users.GetUserByUsername(update.CallbackQuery.From.UserName)
		}
		if update.InlineQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
			user, err = users.GetUserByUsername(update.InlineQuery.From.UserName)
		}
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, err.Error())
			_, _ = bot.Send(msg)
			return true
		}
		println("HandlerId: before runnable", h.id)
		return h.Runnable(bot, h.id, update, user)
	}
	return false
}

func (h *Handler) AsyncLoad(update *tgbotapi.Update) {
	if h.Condition(update) {
		var chatID int64
		var err error
		var user *users.User
		if update.Message != nil {
			chatID = update.Message.Chat.ID
			user, err = users.GetUserByUsername(update.Message.From.UserName)
		}
		if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
			user, err = users.GetUserByUsername(update.CallbackQuery.From.UserName)
		}
		if update.InlineQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
			user, err = users.GetUserByUsername(update.InlineQuery.From.UserName)
		}
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, err.Error())
			_, _ = bot.Send(msg)
			return
		}
		go h.Runnable(bot, h.id, update, user)
	}
}

var handlers []*Handler

func GetHandlers() []*Handler {
	return handlers
}

func AddHandler(condition Condition, runnable Runnable) int {
	id := GeneratorId.GetNext()
	handlers = append(handlers, &Handler{id: id, Condition: condition, Runnable: runnable})
	return id
}

func RemoveHandler(id int) {
	handlers = append(handlers[:id], handlers[id+1:]...)
}

func AddMessageHandler(message string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			return update.Message != nil && update.Message.Text == message
		},
		runnable,
	)
}

func AddCommandHandler(command string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			return update.Message != nil && update.Message.Command() == command
		},
		runnable,
	)
}

func AddQueryDataHandler(data string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			return update.CallbackQuery != nil && update.CallbackData() == data
		},
		runnable,
	)
}
func AddContainsQueryDataHandler(data string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			return update.CallbackQuery != nil && strings.Contains(update.CallbackData(), data)
		},
		runnable,
	)
}

func AddMessageModeHandler(mode string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			user, err := users.GetUserFromUpdate(update)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				_, _ = bot.Send(msg)
				return false
			}
			return update.Message != nil && user.Mode() == mode
		},
		runnable,
	)
}
func AddMessageContainsModeHandler(mode string, runnable Runnable) {
	AddHandler(
		func(update *tgbotapi.Update) bool {
			user, err := users.GetUserFromUpdate(update)
			if err != nil {
				return true
			}
			return update.Message != nil && strings.Contains(user.Mode(), mode)
		},
		runnable,
	)
}
