package repository

import (
	"context"
	"log"
	"online-store/modules/dto"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Add(product dto.ProductDB) error
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Add(product dto.ProductDB) error {

	_, err := r.db.ExecContext(context.Background(),
		`INSERT INTO products(
			product_name, 
			price, 
			stock, 
			image, 
			created_at, 
			created_by, 
			modified_at, 
			modified_by) 
		VALUES(?,?,?,?,?,?,?,?)`,
		product.ProductName,
		product.Price,
		product.Stock,
		product.Image,
		product.CreatedAt,
		product.CreatedBy,
		product.ModifiedAt,
		product.ModifiedBy,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
