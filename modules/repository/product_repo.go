package common

type ProductRepository interface {
}

type productRepository struct {
}

func NewRepo() ProductRepository {
	return &productRepository{}
}
