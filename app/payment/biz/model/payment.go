package model

import (
	"context"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32  `json:"user_id"`
	OrderId       string  `json:"order_id"`
	TranscationId string  `json:"transcation_id"`
	Amount        float64 `json:"amount"`
	PayAt         string  `json:"pay_at"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, paymentLog *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}
