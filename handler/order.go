package handler

import (
	"context"
	"github.com/grpc-shop/order-srv/model"
	"github.com/grpc-shop/order-srv/param"
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

func (o *OrderHandler) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderReply, error) {
	resp := order.CreateOrderReply{
		Code: order.Code_Success,
		Data: new(order.CreateOrderReplyOrder),
	}
	var (
		orderModel model.Order
	)
	orderModel.UserPhone = req.GetUserPhone()
	orderModel.UserName = req.GetUserName()
	orderModel.UserAddress = req.GetUserAddress()
	orderModel.Extra = req.GetExtra()
	orderModel.Remark = req.GetRemark()
	orderModel.OrderStatus = model.WaitPay

	items := req.GetItems()
	for i := range items {
		var (
			itemModel model.OrderItem
		)
		itemModel.ProductId = items[i].GetProductId()
		itemModel.ProductSkuId = items[i].GetProductSkuId()
		itemModel.Price = items[i].GetPrice()
		itemModel.Amount = items[i].GetAmount()
		itemModel.Title = items[i].GetTitle()
		itemModel.Image = items[i].GetImage()
		orderModel.OrderItems = append(orderModel.OrderItems, itemModel)

		orderModel.TotalAmount += itemModel.Amount
	}
	insertId, err := o.server.CreateOrder(&orderModel)
	if err != nil {
		resp.Code = order.Code_CreateErr
		return &resp, err
	}
	resp.Data.Id = insertId
	return &resp, nil
}

func (o *OrderHandler) OrderPaid(ctx context.Context, req *order.OrderPaidReq) (*order.OrderPaidReply, error) {
	resp := order.OrderPaidReply{
		Code: order.Code_Success,
		Data: new(order.OrderPaidReplyOrder),
	}
	createParam := param.InitCreateOrderParam(
		req.GetOrderId(), req.GetPaidAt(), req.GetPaidAmount(), req.GetPaymentNo())
	id, err := o.server.UpdateOrderPay(createParam)
	if err != nil {
		resp.Code = order.Code_OrderPaidErr
		return &resp, err
	}
	resp.Data.Id = id
	return &resp, nil
}
