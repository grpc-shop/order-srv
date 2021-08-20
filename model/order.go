package model

type Order struct {
	Id           int64  `gorm:"column:id" json:"id"`
	No           string `gorm:"column:no" json:"no"`
	UserId       int64  `gorm:"column:user_id" json:"userId"`
	Address      string `gorm:"column:address" json:"address"`
	TotalAmount  int64  `gorm:"column:total_amount" json:"totalAmount"`
	Remark       string `gorm:"column:remark" json:"remark"`
	PaidAt       string `gorm:"column:paid_at" json:"paidAt"`
	PaymentNo    string `gorm:"column:payment_no" json:"paymentNo"`
	RefundStatus string `gorm:"column:refund_status" json:"refundStatus"`
	RefundNo     string `gorm:"column:refund_no" json:"refundNo"`
	OrderStatus  uint8  `gorm:"column:order_status" json:"OrderStatus"`
	Extra        string `gorm:"column:extra" json:"extra"`
	CreatedAt    string `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    string `gorm:"column:updated_at" json:"updatedAt"`
	OrderItems   []OrderItem
}

func (Order) TableName() string {
	return "orders"
}
