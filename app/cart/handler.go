package main

import (
	"context"
	cart "github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
	"github.com/qingz2/gomall/app/cart/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	resp, err = service.NewClearCartService(ctx).Run(req)

	return resp, err
}
