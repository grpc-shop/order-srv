package handler

import (
	"context"
	"github.com/grpc-shop/order-srv/proto/order"
	"github.com/grpc-shop/order-srv/service"
)

var _ order.OrderServer = (*OrderHandler)(nil)

type OrderHandler struct {
	order.UnimplementedOrderServer
	server service.OrderServer
}

func NewOrderHandler(server service.OrderServer) *OrderHandler {
	return &OrderHandler{server: server}
}

func (o OrderHandler) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderReply, error) {
	panic("implement me")
}
