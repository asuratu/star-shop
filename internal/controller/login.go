package controller

import (
	"context"
	"shop/api/v1/backend"
	"shop/internal/service"
	"time"
)

// Login 登录管理
var Login = cLogin{}

type cLogin struct{}

// Login for session
/*func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	err = service.Login().Login(ctx, model.UserLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// for session
	res.Info = service.Session().GetUser(ctx)
	//res.Info = service.BizCtx().Get(ctx).User
	return
}*/

// Login for jwt
func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	var expire time.Time
	res.Token, expire = service.Auth().LoginHandler(ctx)
	res.Expire = expire.Format("2006-01-02 15:04:05")
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
