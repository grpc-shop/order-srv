package param

import "github.com/grpc-shop/order-srv/tool/db"

type PaidOrderParam struct {
	OrderId string
	PaidAt  int64
	Amount  int64
	PayNo   string
}

func InitCreateOrderParam(orderId string, paidAt, amount int64, payNo string) PaidOrderParam {
	return PaidOrderParam{
		OrderId: orderId,
		PaidAt:  paidAt,
		Amount:  amount,
		PayNo:   payNo,
	}
}

type GetListParam struct {
	StartCreateTime int64
	EndCreateTime   int64
	Page            db.Page
}

func InitGetListParam(start, end int64, page, pageSize int64) GetListParam {
	return GetListParam{
		StartCreateTime: start,
		EndCreateTime:   end,
		Page:            db.InitPageSize(page, pageSize),
	}
}
