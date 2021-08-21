package service

import (
	"errors"
	"github.com/grpc-shop/order-srv/dao"
	"github.com/grpc-shop/order-srv/model"
	"github.com/grpc-shop/order-srv/param"
)

type OrderServer interface {
	CreateOrder(orderModel *model.Order) (orderId int64, err error)
	UpdateOrderPay(paidOrderInfo param.PaidOrderParam) (count int64, err error)
}

var _ OrderServer = (*OrderServerImpl)(nil)

var (
	NotFoundOrderErr   = errors.New("订单未找到")
	PaidAmountOrderErr = errors.New("付款金额和订单金额不一致")
	PaidStateErr       = errors.New("订单状态异常")
)

type OrderServerImpl struct {
	orderDao dao.OrderDao
}

func NewOrderServerImpl(dao dao.OrderDao) OrderServer {
	return &OrderServerImpl{orderDao: dao}
}

func (server *OrderServerImpl) CreateOrder(orderModel *model.Order) (orderId int64, err error) {
	return server.orderDao.Create(orderModel)
}

func (server *OrderServerImpl) UpdateOrderPay(paidOrderInfo param.PaidOrderParam) (count int64, err error) {
	order, err := server.orderDao.GetOrder(paidOrderInfo.OrderId)
	if err != nil {
		return 0, NotFoundOrderErr
	}
	if order.TotalAmount != paidOrderInfo.Amount {
		return 0, PaidAmountOrderErr
	}

	if order.OrderStatus != model.WaitPay {
		return 0, PaidStateErr
	}

	return server.orderDao.UpdateOrderPaid(paidOrderInfo)
}
