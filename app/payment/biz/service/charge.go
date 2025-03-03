package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/qingz2/gomall/app/payment/biz/dal/mysql"
	"github.com/qingz2/gomall/app/payment/biz/model"
	payment "github.com/qingz2/gomall/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// validate credit card
	newCard := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	// true: allow test card
	// valid: not allow test card
	err = newCard.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	// create transcation id
	transcationId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	// create payment log
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TranscationId: transcationId.String(),
		Amount:        float64(req.Amount),
		PayAt:         time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	return &payment.ChargeResp{TransactionId: transcationId.String()}, nil
}
