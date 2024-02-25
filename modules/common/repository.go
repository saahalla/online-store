package common

type Repository interface {
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}
