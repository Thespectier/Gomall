package service

import (
	"context"
	"errors"

	"github.com/qingz2/gomall/app/user/biz/dal/mysql"
	"github.com/qingz2/gomall/app/user/biz/module"
	user "github.com/qingz2/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// check required fields
	if req.Email == "" || req.UserName == "" || req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("missing required fields")
	}
	// check password
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password not match")
	}
	// crypt password
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// create user
	newUser := &module.User{
		Username:       req.UserName,
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = module.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	// if everything is ok,
	// return user ID
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
