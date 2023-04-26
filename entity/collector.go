package entity

import "C"
import "errors"

type Collector struct {
	entity
}

type Request struct {
	card   *Card
	owner  *Daimyo
	amount float64
}

var requests []Request

func (c *Collector) AddRequest(number int) error {
	if number >= len(requests) && number <= 0 {
		return errors.New("нет такой заявки")
	}
	request := requests[number]
	requests = append(requests[:number], requests[number+1:]...)
	request.card.balance = request.amount
	return nil
}
