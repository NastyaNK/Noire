package users

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/entity"
)

type User struct {
	entity.IEntity
	mode string
	data map[string]string
}

func (u *User) SetData(key, value string) {
	u.data[key] = value
}

func (u *User) GetData(key string) (string, bool) {
	v, ok := u.data[key]
	return v, ok
}

func (u *User) GetEntity() entity.IEntity {
	return u.IEntity
}

var users []*User
var cards []*entity.Card

func Users() []*User {
	return users
}
func Cards() []*entity.Card {
	return cards
}

func AddCard(card *entity.Card) {
	cards = append(cards, card)
}

func AddUser(user entity.IEntity) {
	users = append(users, &User{IEntity: user, mode: "", data: make(map[string]string)})
}

func GetUserFromUpdate(update *tgbotapi.Update) (*User, error) {
	if update.Message != nil {
		return GetUserByUsername(update.Message.From.UserName)
	}
	if update.CallbackQuery != nil {
		return GetUserByUsername(update.CallbackQuery.From.UserName)
	}
	if update.InlineQuery != nil {
		return GetUserByUsername(update.InlineQuery.From.UserName)
	}
	return nil, errors.New("вы не имеете доступ к функционалу так как вас нет в системе")
}

func (u *User) SetMode(mode string) {
	u.mode = mode
}

func (u *User) Mode() string {
	return u.mode
}

func GetUserByUsername(username string) (*User, error) {
	for _, user := range users {
		if user.Username() == username {
			return user, nil
		}
	}
	return nil, errors.New("такой пользователь не найден")
}
