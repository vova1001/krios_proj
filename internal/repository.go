package inernal

import (
	"database/sql"
)

type PartRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *PartRepo {
	return &PartRepo{db: db}
}
