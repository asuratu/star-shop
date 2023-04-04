package cmd

import (
	"context"
	"shop/internal/consts"
	"shop/internal/controller"
	"shop/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					controller.Hello,        //示例
					controller.Rotation,     // 轮播图
					controller.Position,     // 手工位
					controller.Admin.Create, // 管理员
					controller.Admin.Update, // 管理员
					controller.Admin.Delete, // 管理员
					controller.Login,        // 登录
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					//err := gfAdminToken.Middleware(ctx, group)
					//if err != nil {
					//	panic(err)
					//}
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
						//"/backend/admin/update/me": controller.Admin.Update,
					})
				})
			})
			s.SetSessionMaxAge(time.Minute)
			s.SetSessionStorage(gsession.NewStorageRedisHashTable(g.Redis()))
			s.Run()
			return nil
		},
	}
)
