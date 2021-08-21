package dao

import (
	"github.com/grpc-shop/order-srv/model"
	"github.com/grpc-shop/order-srv/param"
	"gorm.io/gorm"
)

type OrderDao interface {
	Create(order *model.Order) (orderId int64, err error)
	GetOrder(orderId string) (order model.Order, err error)
	UpdateOrderPaid(paidOrderInfo param.PaidOrderParam) (count int64, err error)
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

func (dao *OrderImpl) GetOrder(orderId string) (order model.Order, err error) {
	err = dao.db.Model(&model.Order{}).Scopes(model.GetWithOrderId(orderId)).
		First(&order).Error
	return
}

func (dao *OrderImpl) Create(order *model.Order) (orderId int64, err error) {
	err = dao.db.Model(&model.Order{}).Create(order).Error
	orderId = order.Id
	return
}

func (dao *OrderImpl) UpdateOrderPaid(paidOrderInfo param.PaidOrderParam) (count int64, err error) {
	base := dao.db.Model(&model.Order{}).
		Scopes(model.GetWithOrderId(paidOrderInfo.OrderId)).
		Updates(map[string]interface{}{
			"paid_at":      paidOrderInfo.PaidAt,
			"payment_no":   paidOrderInfo.PayNo,
			"order_status": model.WaitDeliver,
		})
	count = base.RowsAffected
	err = base.Error
	return
}
