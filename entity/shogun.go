package entity

type Shogun struct {
	entity
}

func (*Shogun) GetRank() int {
	return 1
}
