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

func mmsetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 迁移表结构
	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	assert.NoError(t, err)

	return db
}

func TestListOrder_Run(t *testing.T) {
	// 初始化测试数据库
	db := mmsetupTestDB(t)
	mysql.DB = db

	ctx := context.Background()
	s := NewListOrderService(ctx)

	// 测试正常查询订单列表
	t.Run("正常查询订单列表", func(t *testing.T) {
		// 准备测试数据
		order1 := &model.Order{
			OrderId:      "test-order-1",
			UserId:       1,
			UserCurrency: "USD",
			TotalCost:    100.0,
			OrderState:   model.OrderStatePlaced,
			Address: model.Address{
				City:          "Shanghai",
				StreetAddress: "Test Street 1",
				State:         "Test State",
				ZipCode:       200000,
				Email:         "test1@example.com",
			},
		}
		order2 := &model.Order{
			OrderId:      "test-order-2",
			UserId:       1,
			UserCurrency: "USD",
			TotalCost:    200.0,
			OrderState:   model.OrderStatePaid,
			Address: model.Address{
				City:          "Beijing",
				StreetAddress: "Test Street 2",
				State:         "Test State",
				ZipCode:       100000,
				Email:         "test2@example.com",
			},
		}

		// 插入测试数据
		err := db.Create(order1).Error
		assert.NoError(t, err)
		err = db.Create(order2).Error
		assert.NoError(t, err)

		// 添加订单项
		items1 := []*model.OrderItem{
			{
				ProductId:    1,
				OrderIdRefer: "test-order-1",
				Quantity:     2,
				Cost:         50.0,
			},
		}
		items2 := []*model.OrderItem{
			{
				ProductId:    2,
				OrderIdRefer: "test-order-2",
				Quantity:     1,
				Cost:         200.0,
			},
		}

		err = db.Create(items1).Error
		assert.NoError(t, err)
		err = db.Create(items2).Error
		assert.NoError(t, err)

		// 执行测试
		req := &order.ListOrderReq{
			UserId: 1,
		}
		resp, err := s.Run(req)

		// 验证结果
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Orders, 2)

		// 验证订单1
		assert.Equal(t, "test-order-1", resp.Orders[0].OrderId)
		assert.Equal(t, "Shanghai", resp.Orders[0].Address.City)
		assert.Len(t, resp.Orders[0].OrderItems, 1)

		// 验证订单2
		assert.Equal(t, "test-order-2", resp.Orders[1].OrderId)
		assert.Equal(t, "Beijing", resp.Orders[1].Address.City)
		assert.Len(t, resp.Orders[1].OrderItems, 1)
	})

	// 测试查询空结果
	t.Run("查询空结果", func(t *testing.T) {
		req := &order.ListOrderReq{
			UserId: 999, // 使用不存在的用户ID
		}
		resp, err := s.Run(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.Orders)
	})
}
