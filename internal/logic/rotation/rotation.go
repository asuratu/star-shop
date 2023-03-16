package content

import (
	"context"
	"shop/internal/consts"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out *model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.RotationCreateOutput{RotationId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Unscoped().Delete() // 加上 Unscoped() 代表真正的删除
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		columns := dao.RotationInfo.Columns()
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(columns.Id).
			Where(columns.Id, in.Id).
			Update()
		return err
	})
}

// Index 分页获取列表
func (s *sRotation) Index(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	var (
		m = dao.RotationInfo.Ctx(ctx)
	)
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	switch in.Sort {
	case consts.RotationSortByCreateTimeDesc:
		listModel = listModel.OrderDesc(dao.RotationInfo.Columns().CreatedAt)
	case consts.RotationSortByUpdateTimeDesc:
		listModel = listModel.OrderDesc(dao.RotationInfo.Columns().UpdatedAt)
	default:
		listModel = listModel.OrderDesc(dao.RotationInfo.Columns().Id)
	}
	// 执行查询
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	// 没有数据
	if len(out.List) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	return out, nil
}
