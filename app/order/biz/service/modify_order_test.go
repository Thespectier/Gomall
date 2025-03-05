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

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	assert.NoError(t, err)

	mysql.DB = db
	return db
}

func TestModifyOrder_Run(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()
	s := NewModifyOrderService(ctx)

	// 测试正常修改订单
	t.Run("正常修改订单", func(t *testing.T) {
		// 准备测试数据
		testOrder := &model.Order{
			OrderId:      "test-order-1",
			UserId:       1,
			UserCurrency: "USD",
			TotalCost:    100.0,
			OrderState:   model.OrderStatePlaced,
			Address: model.Address{
				City:          "Shanghai",
				StreetAddress: "Test Street",
				State:         "Test State",
				ZipCode:       200000,
				Email:         "test@example.com",
			},
		}

		// 插入测试数据
		err := db.Create(testOrder).Error
		assert.NoError(t, err)

		// 添加原始订单项
		originalItems := []*model.OrderItem{
			{
				ProductId:    1,
				OrderIdRefer: "test-order-1",
				Quantity:     1,
				Cost:         100.0,
			},
		}
		err = db.Create(originalItems).Error
		assert.NoError(t, err)

		// 执行测试
		req := &order.ModifyOrderReq{
			UserId:       1,
			OrderId:      "test-order-1",
			UserCurrency: "USD",
			Address: &order.Address{
				City:          "Beijing",
				StreetAddress: "New Street",
				State:         "New State",
				ZipCode:       100000,
				Country:       "China",
			},
			Email: "new@example.com",
			OrderItems: []*order.OrderItem{
				{
					Item: &order.CartItem{
						ProductId: 2,
						Quantity:  2,
					},
					Cost: 150.0,
				},
			},
		}
		resp, err := s.Run(req)

		// 验证结果
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "test-order-1", resp.OrderId)

		// 验证订单是否已更新
		var updatedOrder model.Order
		err = db.First(&updatedOrder, "order_id = ?", "test-order-1").Error
		assert.NoError(t, err)
		assert.Equal(t, model.OrderStateModified, updatedOrder.OrderState)
		assert.Equal(t, "Beijing", updatedOrder.Address.City)
		assert.Equal(t, "New Street", updatedOrder.Address.StreetAddress)
		assert.Equal(t, float32(150.0), updatedOrder.TotalCost)

		// 验证订单项是否已更新
		var updatedItems []*model.OrderItem
		err = db.Where("order_id_refer = ?", "test-order-1").Find(&updatedItems).Error
		assert.NoError(t, err)
		assert.Len(t, updatedItems, 1)
		assert.Equal(t, uint32(2), updatedItems[0].ProductId)
		assert.Equal(t, int32(2), updatedItems[0].Quantity)
		assert.Equal(t, float32(150.0), updatedItems[0].Cost)
	})

	// 测试无效订单ID
	t.Run("无效订单ID", func(t *testing.T) {
		req := &order.ModifyOrderReq{
			UserId:       1,
			OrderId:      "non-existent-order",
			UserCurrency: "USD",
			Address: &order.Address{
				City:          "Beijing",
				StreetAddress: "New Street",
				State:         "New State",
				ZipCode:       100000,
				Country:       "China",
			},
			Email: "new@example.com",
			OrderItems: []*order.OrderItem{
				{
					Item: &order.CartItem{
						ProductId: 2,
						Quantity:  2,
					},
					Cost: 150.0,
				},
			},
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "order not found")
	})

	// 测试无效请求参数
	t.Run("无效请求参数", func(t *testing.T) {
		req := &order.ModifyOrderReq{
			UserId:       0, // 无效的用户ID
			OrderId:      "test-order-1",
			UserCurrency: "",
			Address:      nil,
			Email:        "",
			OrderItems:   []*order.OrderItem{},
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
