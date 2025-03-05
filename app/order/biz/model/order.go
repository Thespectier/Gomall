package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Address 地址结构体
type Address struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

// OrderState 订单状态
type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"   // 已下单
	OrderStatePaid     OrderState = "paid"     // 已支付
	OrderStateShipped  OrderState = "shipped"  // 已发货
	OrderStateCanceled OrderState = "canceled" // 已取消
	OrderStateModified OrderState = "modified" // 已修改
)

// Order 订单结构体
type Order struct {
	Base
	PaidAt         time.Time
	AutoCanceledAt time.Time
	OrderId        string `gorm:"uniqueIndex;size:256"`
	UserId         uint32
	UserCurrency   string
	Address        Address     `gorm:"embedded"`
	OrderItems     []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState     OrderState
	TotalCost      float32
}

// TableName 表名
func (o Order) TableName() string {
	return "order"
}

// CreateOrder 创建订单
func CreateOrder(db *gorm.DB, ctx context.Context, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

// 获取订单列表
func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

// 获取订单详情
func GetOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (order Order, err error) {
	err = db.WithContext(ctx).Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

// 修改订单状态
func UpdateOrderState(db *gorm.DB, ctx context.Context, userId uint32, orderId string, state OrderState) error {
	return db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
}

// 修改订单
func UpdateOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string, updates map[string]interface{}) error {
	return db.WithContext(ctx).Model(&Order{}).
		Where(&Order{UserId: userId, OrderId: orderId}).
		Updates(updates).Error
}
