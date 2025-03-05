package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type ModifyOrderService struct {
	ctx context.Context
} // NewModifyOrderService new ModifyOrderService
func NewModifyOrderService(ctx context.Context) *ModifyOrderService {
	return &ModifyOrderService{ctx: ctx}
}

// Run create note info
func (s *ModifyOrderService) Run(req *order.ModifyOrderReq) (resp *order.ModifyOrderResp, err error) {
	// Finish your business logic.
	// 判定输入是否合法
	if req.UserId <= 0 || req.OrderId == "" || req.UserCurrency == "" || req.Address == nil || req.Email == "" || len(req.OrderItems) == 0 {
		err = fmt.Errorf("invalid input")
		return
	}
	// 开启事务保持原子性
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//更新订单商品项和总价
		var itemlist []*model.OrderItem
		var totalCost float32
		for _, item := range req.OrderItems {
			itemlist = append(itemlist, &model.OrderItem{
				ProductId:    item.Item.ProductId,
				OrderIdRefer: req.OrderId,
				Quantity:     item.Item.Quantity,
				Cost:         item.Cost,
			})
			totalCost += item.Cost
		}
		// 更新订单地址信息
		address := model.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			ZipCode:       req.Address.ZipCode,
			Email:         req.Email,
			Country:       req.Address.Country,
		}
		// 创建待更新项map
		updateMap := map[string]interface{}{
			"OrderStata":   model.OrderStateModified,
			"UserCurrency": req.UserCurrency,
			"Address":      address,
			"OrderItems":   itemlist, //此处可能存在问题，待修改
			"TotalCost":    totalCost,
		}
		// 更新订单信息
		if err := model.UpdateOrder(mysql.DB, s.ctx, req.UserId, req.OrderId, updateMap); err != nil {
			klog.Errorf("model.UpdateOrder error: %v", err)
			return err
		}
		// 返回响应
		resp = &order.ModifyOrderResp{
			OrderId: req.OrderId,
		}
		return nil
	})
	if err != nil {
		klog.Errorf("mysql.DB.Transaction error: %v", err)
		return
	}
	return
}
