//+build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/grpc-shop/order-srv/dao"
	"github.com/grpc-shop/order-srv/service"
	"gorm.io/gorm"
)

func InitOrderHandler(db *gorm.DB) *OrderHandler {
	panic(wire.Build(
		dao.NewOrderImpl,
		service.NewOrderServerImpl,
		NewOrderHandler,
	))
	return &OrderHandler{}
}
