package login

import (
	"context"
	"shop/internal/consts"
	"shop/internal/service"
	"shop/utility/xerr"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	utility "shop/utility/utils"

	"github.com/gogf/gf/v2/util/gutil"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// Login 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	gutil.Dump("加密后密码：", utility.EncryptPassword(in.Password, adminInfo.UserSalt))
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return xerr.WrapErr(consts.AdminNameOrPasswordError)
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线 for session
	//service.BizCtx().SetUser(ctx, &model.ContextUser{
	//	Id:      uint(adminInfo.Id),
	//	Name:    adminInfo.Name,
	//	IsAdmin: uint8(adminInfo.IsAdmin),
	//})
	return nil
}

// Logout 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
