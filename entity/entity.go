package entity

var (
	SHOGUN    = "shogun"
	DAIMYO    = "daimyo"
	SAMURAI   = "samurai"
	COLLECTOR = "collector"
)

type entity struct {
	username string
	nickname string
	entities []IEntity
}

type IEntity interface {
	Username() string
	SetName(string)
	Nickname() string
	SetUsername(string)
	addSubordinate(IEntity)
	Subordinates() []IEntity
	GetRank() int
}

func (e *entity) Username() string {
	return e.username
}

func (e *entity) SetName(name string) {
	e.username = name
}

func (e *entity) Nickname() string {
	return e.nickname
}

func (e *entity) SetUsername(username string) {
	e.nickname = username
}

func (e *entity) Subordinates() []IEntity {
	return e.entities
}

func (e *entity) GetRank() int {
	return 0
}

func (e *entity) addSubordinate(entity IEntity) {
	e.entities = append(e.entities, entity)
}
