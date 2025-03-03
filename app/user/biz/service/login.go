package service

import (
	"context"
	"errors"

	"github.com/qingz2/gomall/app/user/biz/dal/mysql"
	"github.com/qingz2/gomall/app/user/biz/module"
	user "github.com/qingz2/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// check username and password
	if req.UserName == "" || req.Password == "" {
		return nil, errors.New("username or password is empty")
	}
	// get user by username
	row, err := module.GetByUsername(mysql.DB, req.UserName)
	if err != nil {
		return nil, err
	}
	// check password
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	// if everything is ok,
	// return user ID
	resp = &user.LoginResp{UserId: int32(row.ID)}

	return resp, nil
}
