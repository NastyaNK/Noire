package entity

import (
	"errors"
)

type Admin struct {
	entity
}

func NewAdmin(username, nickname string) *Admin {
	return &Admin{entity: entity{
		username: username,
		nickname: nickname,
		entities: make([]IEntity, 0),
	}}
}

func (a *Admin) NewEntity(entityName, username, nickname string) (IEntity, error) {
	println(entityName, username, nickname)
	e := entity{
		username: username,
		nickname: nickname,
		entities: make([]IEntity, 0),
	}
	switch entityName {
	case SHOGUN:
		return &Shogun{e}, nil
	case DAIMYO:
		return &Daimyo{e, nil}, nil
	case SAMURAI:
		return &Samurai{e}, nil
	case COLLECTOR:
		return &Collector{e}, nil
	}
	return nil, errors.New("такого типа сущности нет")
}

func (a *Admin) Assign(e1, e2 IEntity) error {
	if e2.GetRank()-1 == e1.GetRank() {
		e1.addSubordinate(e2)
	} else {
		return errors.New("")
	}
	return nil
}
