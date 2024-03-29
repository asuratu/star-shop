package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string `json:"pic_url"`
	Link   string `json:"link"`
	Sort   int    `json:"sort"`
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint `json:"id"`
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型，1:按照创建时间排序 2:按照更新时间排序
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationListItem `json:"list" description:"列表"`
	Page  int                `json:"page" description:"分页码"`
	Size  int                `json:"size" description:"分页数量"`
	Total int                `json:"total" description:"数据总数"`
}

// RotationListItem 主要用于列表展示
type RotationListItem struct {
	Id        uint        `json:"id"`         // 自增ID
	PicUrl    string      `json:"pic_url"`    // 图片地址
	Link      string      `json:"link"`       // 链接地址
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
