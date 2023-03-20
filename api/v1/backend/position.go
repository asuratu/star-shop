package backend

import (
	v1 "shop/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

type PositionAddReq struct {
	g.Meta    `path:"/positions" tags:"Position" method:"post" summary:"You first position api"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
}

type PositionAddRes struct {
	PositionId uint64 `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/positions/{id}" method:"delete" tags:"手工位图" summary:"删除手工位图接口"`
}
type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/positions/{id}" method:"put" tags:"手工位图" summary:"修改手工位图接口"`
	PicUrl    string `json:"pic_url" v:"required#手工位图图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort      int    `json:"sort"    dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
}
type PositionUpdateRes struct{}

type PositionGetListCommonReq struct {
	g.Meta `path:"/positions" method:"get" tags:"手工位图" summary:"手工位图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	v1.CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
