package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	// 从数据库中获取订单列表
	orders, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("ListOrder failed, err: %v", err)
		return nil, err
	}
	// 进行数据映射
	var orderList []*order.Order
	for _, w := range orders {
		//先进行订单商品列表的映射
		var orderItemList []*order.OrderItem
		for _, item := range w.OrderItems {
			orderItemList = append(orderItemList, &order.OrderItem{
				Item: &order.CartItem{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				},
				Cost: item.Cost,
			},
			)
		}
		// 再进行订单信息的映射
		tmp := &order.Order{
			OrderId:      w.OrderId,
			UserId:       w.UserId,
			UserCurrency: w.UserCurrency,
			Email:        w.Address.Email,
			Address: &order.Address{
				StreetAddress: w.Address.StreetAddress,
				City:          w.Address.City,
				State:         w.Address.State,
				ZipCode:       int32(w.Address.ZipCode),
			},
			OrderItems:   orderItemList,
			CreatedAt:    int64(w.CreatedAt.Unix()),
			PaidAt:       int64(w.PaidAt.Unix()),
			UpdatedAt:    int64(w.UpdatedAt.Unix()),
			AutocancelAt: int64(w.AutoCanceledAt.Unix()),
		}
		orderList = append(orderList, tmp)
	}
	// 组装响应数据
	resp = &order.ListOrderResp{
		Orders: orderList,
	}

	return
}
