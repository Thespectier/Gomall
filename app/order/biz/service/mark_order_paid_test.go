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

func initTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	assert.NoError(t, err)

	mysql.DB = db
	return db
}

func createTestOrder(db *gorm.DB, orderId string, userId uint32, state model.OrderState) {
	testOrder := &model.Order{
		OrderId:      orderId,
		UserId:       userId,
		UserCurrency: "USD",
		TotalCost:    100.0,
		OrderState:   state,
		Address: model.Address{
			City:          "Shanghai",
			StreetAddress: "Test Street",
			State:         "Test State",
			ZipCode:       200000,
			Email:         "test@example.com",
		},
	}
	err := db.Create(testOrder).Error
	if err != nil {
		panic(err)
	}
}

func TestMarkOrderPaid_Run(t *testing.T) {
	db := initTestDB(t)
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)

	// 测试正常支付订单
	t.Run("正常支付订单", func(t *testing.T) {
		createTestOrder(db, "test-order-1", 1, model.OrderStatePlaced)

		req := &order.MarkOrderPaidReq{
			UserId:  1,
			OrderId: "test-order-1",
		}
		resp, err := s.Run(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)

		var updatedOrder model.Order
		err = db.First(&updatedOrder, "order_id = ?", "test-order-1").Error
		assert.NoError(t, err)
		assert.Equal(t, model.OrderStatePaid, updatedOrder.OrderState)
	})

	// 测试无效订单ID
	t.Run("无效订单ID", func(t *testing.T) {
		req := &order.MarkOrderPaidReq{
			UserId:  1,
			OrderId: "non-existent-order",
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// 测试无效用户ID
	t.Run("无效用户ID", func(t *testing.T) {
		req := &order.MarkOrderPaidReq{
			UserId:  0,
			OrderId: "test-order-1",
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// 测试已支付订单
	t.Run("已支付订单", func(t *testing.T) {
		createTestOrder(db, "test-order-2", 1, model.OrderStatePaid)

		req := &order.MarkOrderPaidReq{
			UserId:  1,
			OrderId: "test-order-2",
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// 测试已取消订单
	t.Run("已取消订单", func(t *testing.T) {
		createTestOrder(db, "test-order-3", 1, model.OrderStateModified)

		req := &order.MarkOrderPaidReq{
			UserId:  1,
			OrderId: "test-order-3",
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// 测试用户ID不匹配
	t.Run("用户ID不匹配", func(t *testing.T) {
		createTestOrder(db, "test-order-4", 1, model.OrderStatePlaced)

		req := &order.MarkOrderPaidReq{
			UserId:  2,
			OrderId: "test-order-4",
		}
		resp, err := s.Run(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
