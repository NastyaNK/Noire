package entity

import (
	"errors"
)

type Daimyo struct {
	entity
	cards []*Card
}

func (d *Daimyo) Cards() []*Card {
	return d.cards
}

func (d *Daimyo) AddCard(card *Card) {
	d.cards = append(d.cards, card)
}

func (*Daimyo) GetRank() int {
	return 2
}
func (d *Daimyo) AddRequest(cardNumber int, amount float64) error {
	for i := 0; i < len(d.cards); i++ {
		if d.cards[i].number == cardNumber {
			requests = append(requests, Request{d.cards[i], d, amount})
			return nil
		}
	}
	return errors.New("нет такой карты")
}
