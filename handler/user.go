package handler

import (
	"context"
	"github.com/ashoreDove/common"
	"github.com/ashoreDove/parasite-user/domain/model"
	"github.com/ashoreDove/parasite-user/domain/service"
	user "github.com/ashoreDove/parasite-user/proto/user"
	"github.com/jinzhu/gorm"
)

type IUserService interface {
	//注册
	Register(context.Context, *user.RegisterRequest, *user.Response) error
	//登录
	Login(context.Context, *user.LoginRequest, *user.LoginResponse) error
	//验证码的短信发送
	SendMessage(context.Context, *user.MessageRequest, *user.MessageResponse) error
	//todo 验证码的校验
}

type User struct {
	UserDataService service.IUserDataService
}

func (u User) SendMessage(ctx context.Context, request *user.MessageRequest, response *user.MessageResponse) error {
	//TODO implement me
	panic("implement me")
}

func (u User) Register(ctx context.Context, req *user.RegisterRequest, resp *user.Response) error {
	//TODO 手机验证码校验

	newUser := &model.User{
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Nickname,
	}
	_, err := u.UserDataService.AddUser(newUser)
	if err != nil {
		resp.Msg = "Failed"
		return err
	}
	resp.Msg = "Successfully"
	return nil
}

func (u User) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) error {
	//TODO 手机验证码校验
	isOk, err := u.UserDataService.CheckPassword(req.Account, req.Password)
	if isOk {
		resp.Msg = "Successfully"
		//token
		token, err := common.GenerateToken(req.Account, req.Password)
		if err != nil {
			resp.Msg = "Token Create Failed"
			return err
		}
		resp.Token = token
	}
	resp.Msg = "Failed"
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(db *gorm.DB) IUserService {
	return &User{service.NewUserDataService(db)}
}
