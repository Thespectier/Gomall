package service

import (
	"context"

	"github.com/qingz2/gomall/app/product/biz/dal/mysql"
	"github.com/qingz2/gomall/app/product/biz/model"
	product "github.com/qingz2/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	if mysql.DB == nil {
		return nil, err
	}

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.SearchProduct(req.Query)
	if err != nil {
		return nil, err
	}

	var results []*product.Product
	for _, v := range p {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.SearchProductsResp{Results: results}, err
}
