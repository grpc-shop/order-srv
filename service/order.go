package service

import (
	"github.com/grpc-shop/order-srv/dao"
)

type OrderServer interface {
}

var _ OrderServer = (*OrderServerImpl)(nil)

type OrderServerImpl struct {
	orderDao dao.OrderDao
}

func NewOrderServerImpl(dao dao.OrderDao) OrderServer {
	return &OrderServerImpl{orderDao: dao}
}
