package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/qingz2/gomall/app/cart/biz/dal/mysql"
	cart "github.com/qingz2/gomall/rpc_gen/kitex_gen/cart"
)

func TestGetCart_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Logf("err: %v", err)
	}
	mysql.Init()
	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	req := &cart.GetCartReq{
		UserId: 22,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
