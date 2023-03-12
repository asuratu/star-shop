package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Api 定义业务侧的数据结构, 提供对外接口的输入|输出参数定义

type RotationAddReq struct {
	g.Meta `path:"/rotations" tags:"Rotation" method:"post" summary:"add rotation api"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}

type RotationAddRes struct {
	RotationId uint `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta     `path:"/rotations" tags:"Rotation" method:"delete" summary:"delete rotation api"`
	RotationId uint `json:"rotation_id" v:"required|min:1#轮播图ID不能为空|轮播图ID最小为1" dc:"轮播图ID"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotations/{id}" tags:"Rotation" method:"put" summary:"update rotation api"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}

type RotationUpdateRes struct{}
