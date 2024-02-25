package dto

import "time"

type DefaultDate struct {
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string    `db:"updated_by"`
	DeletedAt time.Time `db:"deleted_at"`
	DeletedBy string    `db:"deleted_by"`
}
