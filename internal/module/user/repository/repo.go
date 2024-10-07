package repository

import (
	"github.com/jmoiron/sqlx"
)

var _ ports.XxxRepository = &xxxRepository{}

type xxxRepository struct {
	db *sqlx.DB
}

func NewXxxRepository(db *sqlx.DB) *xxxRepository {
	return &xxxRepository{
		db: db,
	}
}
