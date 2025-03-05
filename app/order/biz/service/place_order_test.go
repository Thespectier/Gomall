package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func psetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 迁移表结构
	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	assert.NoError(t, err)

	return db
}

func TestPlaceOrder_Run(t *testing.T) {
	// 初始化测试数据库
	db := psetupTestDB(t)
	mysql.DB = db

	ctx := context.Background()
	s := NewPlaceOrderService(ctx)

	// 测试正常下单
	t.Run("正常下单", func(t *testing.T) {
		req := &order.PlaceOrderReq{
			UserId:       1,
			UserCurrency: "USD",
			Address: &order.Address{
				City:          "Shanghai",
				StreetAddress: "Test Street",
				State:         "Test State",
				ZipCode:       200000,
			},
			Email:      "xxx@xxx.xxx",
			OrderItems: []*order.OrderItem{{Item: &order.CartItem{ProductId: 1, Quantity: 1}, Cost: 10.0}},
		}
		resp, err := s.Run(req)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.OrderId)

		// 验证订单是否成功创建
		var order model.Order
		err = db.First(&order, "order_id = ?", resp.OrderId).Error
		assert.NoError(t, err)
		assert.Equal(t, float32(10.0), order.TotalCost)
	})

	// 测试空订单项
	t.Run("空订单项", func(t *testing.T) {
		req := &order.PlaceOrderReq{
			UserId:       1,
			UserCurrency: "USD",
			OrderItems:   []*order.OrderItem{},
		}
		resp, err := s.Run(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
