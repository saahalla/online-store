package repository

import "github.com/jmoiron/sqlx"

type AuthRepository interface {
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

// func (r *authRepository)
