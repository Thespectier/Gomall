package service

import (
	"context"
	"testing"

	//"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value
	req := &order.PlaceOrderReq{
		UserId:       1,
		UserCurrency: "USD",
		Address:      &order.Address{City: "Shanghai"},
		Email:        "xxx@xxx.xxx",
		OrderItems:   []*order.OrderItem{{Item: &order.CartItem{ProductId: 1, Quantity: 1}, Cost: 10.0}},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
