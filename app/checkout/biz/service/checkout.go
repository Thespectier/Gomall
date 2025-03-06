package service

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/qingz2/gomall/app/checkout/infra/mq"
	checkout "github.com/qingz2/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/qingz2/gomall/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          "req.Email",
		ContentType: "text/plain",
		Subject:     "You have just created an order in our shop",
		Content:     "You have just created an order in our shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)
	return
}
