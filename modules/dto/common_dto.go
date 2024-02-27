package dto

import "time"

type DefaultDate struct {
	CreatedAt  time.Time `db:"created_at" goqu:"skipupdate"`
	CreatedBy  string    `db:"created_by" goqu:"skipupdate"`
	ModifiedAt time.Time `db:"modified_at"`
	ModifiedBy string    `db:"modified_by"`
}
