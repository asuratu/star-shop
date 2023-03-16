package controller

import (
	"context"
	v1 "shop/api/v1"
	"shop/api/v1/backend"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationAddReq) (res *backend.RotationAddRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationAddRes{
		RotationId: out.RotationId,
	}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.RotationId)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	// 获取 path 中的 id
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		// 获取 path 中的 id
		Id: ghttp.RequestFromCtx(ctx).Get("id").Uint(),
	})
	if err != nil {
		return nil, err
	}
	return
}

func (a *cRotation) Index(ctx context.Context, req *backend.RotationGetListReq) (res *backend.RotationGetListRes, err error) {
	out, err := service.Rotation().Index(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationGetListRes{
		CommonPaginationRes: v1.CommonPaginationRes{
			Total: out.Total,
			Page:  out.Page,
			Size:  out.Size,
		},
		List: out.List,
	}, nil
}
