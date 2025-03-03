package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/qingz2/gomall/app/cart/biz/dal/model"
	"github.com/qingz2/gomall/app/cart/biz/dal/mysql"
	cart "github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// check product exist
	/*
		rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp == nil || productResp.Product.id == 0 {
			return nil, kerrors.NewBizStatusError(40004, "product not exist")
		}
	*/
	// add item
	cartItem := model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}

	err = model.AddItem(s.ctx, mysql.DB, &cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
