package controller

import (
	"context"
	"shop/api/v1/backend"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionAddReq) (res *backend.PositionAddRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionAddRes{
		PositionId: out.PositionId,
	}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	Id := ghttp.RequestFromCtx(ctx).Get("id").Uint64()
	g.Dump(Id)
	err = service.Position().Delete(ctx, Id)
	if err != nil {
		return nil, err
	}
	return &backend.PositionDeleteRes{}, nil
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
		Id: ghttp.RequestFromCtx(ctx).Get("id").Uint64(),
	})
	return &backend.PositionUpdateRes{}, nil
}

func (a *cPosition) Index(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionGetListCommonRes{
		Total: getListRes.Total,
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
	}, nil
}
