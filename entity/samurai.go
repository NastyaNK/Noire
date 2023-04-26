package entity

type Samurai struct {
	entity
}

func (*Samurai) GetRank() int {
	return 3
}
