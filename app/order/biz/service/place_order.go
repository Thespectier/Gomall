package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// 判断订单商品项是否合法
	if len(req.OrderItems) <= 0 {
		klog.Errorf("order items is empty")
		err = fmt.Errorf("OrderItems empty")
		return
	}
	// 开启事务保持原子性
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 生成订单id
		orderId, err := uuid.NewUUID()
		if err != nil {
			klog.Errorf("generate order id error: %v", err)
			return err
		}
		// 保存订单信息
		tmp := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			OrderState:   model.OrderStatePlaced,
		}
		// 保存订单地址信息
		if req.Address != nil {
			tmp.Address.City = req.Address.City
			tmp.Address.Email = req.Email
			tmp.Address.StreetAddress = req.Address.StreetAddress
			tmp.Address.State = req.Address.State
			tmp.Address.ZipCode = req.Address.ZipCode
		}
		// 保存商品项信息
		var itemlist []*model.OrderItem
		var totalCost float32
		for _, item := range req.OrderItems {
			itemlist = append(itemlist, &model.OrderItem{
				ProductId:    item.Item.ProductId,
				OrderIdRefer: orderId.String(),
				Quantity:     item.Item.Quantity,
				Cost:         item.Cost,
			})
			totalCost += item.Cost
		}
		// 保存订单总价
		tmp.TotalCost = totalCost
		// 订单插入数据库
		if err := tx.Create(tmp).Error; err != nil {
			klog.Errorf("create order error: %v", err)
			return err
		}
		// 订单商品项插入数据库
		if err := tx.Create(&itemlist).Error; err != nil {
			klog.Errorf("create order item error: %v", err)
			return err
		}
		// 创建返回消息
		resp = &order.PlaceOrderResp{
			OrderId: orderId.String(),
		}
		return nil
	})
	if err != nil {
		klog.Errorf("place order error: %v", err)
		return
	}
	return
}

