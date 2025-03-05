package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	// 判定传入数据是否合法
	if req.UserId <= 0 || req.OrderId == "" {
		err = fmt.Errorf("invalid request: user_id or order_id is empty")
		klog.Error(err)
		return
	}
	// 判断订单是否存在
	_, err = model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.GetOrder error: %v", err)
		return nil, err
	}
	// 更新订单状态为已支付
	err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.UpdateOrderState error: %v", err)
		return nil, err
	}
	// 返回响应
	resp = &order.MarkOrderPaidResp{}

	return
}
