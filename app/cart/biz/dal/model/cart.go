package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id;comment:用户ID"`
	ProductId uint32 `gorm:"type:int(11);not null;comment:商品ID"`
	Qty       uint32 `gorm:"type:int(11);not null;comment:商品数量"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	// check if product exist in cart
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// if product exist, update qty
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	// if product not exist, create new record
	return db.WithContext(ctx).Create(item).Error
}

func ClearCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	// check if user exist
	if userId == 0 {
		return errors.New("invalid user id")
	}
	// delete all items in cart
	return db.WithContext(ctx).
		Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCart(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: userId}).
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
