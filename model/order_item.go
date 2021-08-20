package model

type OrderItem struct {
	Id           int64  `gorm:"column:id" json:"id"`
	OrderId      int64  `gorm:"column:order_id" json:"orderId"`
	ProductId    int64  `gorm:"column:product_id" json:"productId"`
	ProductSkuId int64  `gorm:"column:product_sku_id" json:"productSkuId"`
	Amount       int64  `gorm:"column:amount" json:"amount"`
	Price        int64  `gorm:"column:price" json:"price"`
	Title        string `gorm:"column:title" json:"title"`
	Image        string `gorm:"column:image" json:"image"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
