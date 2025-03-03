package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/qingz2/gomall/app/cart/biz/dal/model"
	"github.com/qingz2/gomall/app/cart/biz/dal/mysql"
	cart "github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// get cart by user id
	itemList, err := model.GetCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	// convert to rpc struct
	var items []*cart.CartItem
	for _, item := range itemList {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Qty),
		})
	}
	getCart := cart.Cart{UserId: req.UserId, Items: items}
	return &cart.GetCartResp{Cart: &getCart}, nil
}
