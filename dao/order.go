package dao

import (
	"gorm.io/gorm"
)

type OrderDao interface {
}

var _ OrderDao = (*OrderImpl)(nil)

type OrderImpl struct {
	db *gorm.DB
}

func NewOrderImpl(db *gorm.DB) OrderDao {
	return &OrderImpl{
		db: db,
	}
}
