package entity

import (
	"errors"
)

type Bank struct {
	info string
}

func NewBank(info string) *Bank {
	return &Bank{info: info}
}

func (b *Bank) Info() string {
	return b.info
}

type Card struct {
	bank    *Bank
	limit   int
	number  int
	balance float64
}

func NewCard(ie IEntity, bank *Bank) (*Card, error) {
	switch ie.(type) {
	case *Admin, *Shogun:
		return &Card{
			bank:  bank,
			limit: 200000,
		}, nil
	}
	return nil, errors.New("вы не можете создават карту")
}

func AssignCard(e1, e2 IEntity, c *Card) error {
	switch e1.(type) {
	case *Admin, *Shogun:
		d, ok := e2.(*Daimyo)
		if ok {
			d.AddCard(c)
			return nil
		}
		return errors.New("")
	}
	return errors.New("")
}
