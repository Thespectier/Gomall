package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/qingz2/gomall/app/cart/biz/dal/model"
	"github.com/qingz2/gomall/app/cart/biz/dal/mysql"
	cart "github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
)

type ClearCartService struct {
	ctx context.Context
} // NewClearCartService new ClearCartService
func NewClearCartService(ctx context.Context) *ClearCartService {
	return &ClearCartService{ctx: ctx}
}

// Run create note info
func (s *ClearCartService) Run(req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	// clear cart
	err = model.ClearCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.ClearCartResp{}, nil
}
