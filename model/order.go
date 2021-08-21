package model

import (
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type OrderState int16

const (
	WaitPay     OrderState = iota + 1 //待付款
	WaitDeliver                       //待发货
	ForGood                           //待收货
	Cancel      = 10                  // 已取消
)

type Order struct {
	Id           int64       `gorm:"column:id" json:"id"`
	OrderId      string      `gorm:"column:order_id" json:"orderId"`
	UserId       int64       `gorm:"column:user_id" json:"userId"`
	UserName     string      `gorm:"column:user_name" json:"userName"`
	UserPhone    string      `gorm:"column:user_phone" json:"userPhone"`
	UserAddress  string      `gorm:"column:user_address" json:"userAddress"`
	TotalAmount  int64       `gorm:"column:total_amount" json:"totalAmount"`
	Remark       string      `gorm:"column:remark" json:"remark"`
	PaidAt       int64       `gorm:"column:paid_at" json:"paidAt"`
	PaymentNo    string      `gorm:"column:payment_no" json:"paymentNo"`
	RefundStatus string      `gorm:"column:refund_status" json:"refundStatus"`
	RefundNo     string      `gorm:"column:refund_no" json:"refundNo"`
	OrderStatus  OrderState  `gorm:"column:order_status" json:"OrderStatus"`
	Extra        string      `gorm:"column:extra" json:"extra"`
	CreatedAt    int64       `gorm:"column:created_at;autoCreatedAt" json:"createdAt"`
	UpdatedAt    int64       `gorm:"column:updated_at;autoUpdatedAt" json:"updatedAt"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderId;references:OrderId"`
}

func (Order) TableName() string {
	return "orders"
}

func GetWithOrderId(orderId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("order_id=?", orderId)
	}
}
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.OrderId == "" {
		o.OrderId = createOrderId()
	}
	return nil
}

// 时间戳+7位随机
func createOrderId() string {
	var (
		builder strings.Builder
	)
	builder.WriteString(strconv.Itoa(int(time.Now().Unix())))
	for i := 0; i < 7; i++ {
		builder.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return builder.String()
}
