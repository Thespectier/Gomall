package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/qingz2/gomall/app/checkout/infra/rpc"
	"github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/qingz2/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/qingz2/gomall/rpc_gen/kitex_gen/payment"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// get cart info
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		return nil, kerrors.NewGRPCBizStatusError(5005002, "cart is empty")
	}
	// get total product price in cart
	/*
		var total float32
		for _, cartItem := range cartResult.Cart.Items {
			productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{ProductId: cartItem.ProductId})
			if resultErr != nil {
				return nil, resultErr
			}
			if productResp.Product == nil {
				continue
			}
			p := productResp.Product.Price
			cost := p * float32(cartItem.Quantity)
			total += cost
		}
	*/
	// create order id
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()
	// create payment request
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		//Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}
	// clear cart
	_, err = rpc.CartClient.ClearCart(s.ctx, &cart.ClearCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	// charge the payment
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	// save order info
	klog.Info(paymentResult)
	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return resp, nil
}
