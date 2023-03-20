package position

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// Create 创建
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out *model.PositionCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.PositionCreateOutput{
		PositionId: uint64(lastInsertID),
	}, err
}

// Delete 删除
func (s *sPosition) Delete(ctx context.Context, id uint64) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: id,
		}).Delete() // 加上 Unscoped() 代表真正的删除, 否则是软删除
		return err
	})
}

// Update 修改
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		columns := dao.PositionInfo.Columns()
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(columns.Id).
			Where(columns.Id, in.Id).
			Update()
		return err
	})
}

// GetList 获取
func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error) {
	var (
		m = dao.PositionInfo.Ctx(ctx)
	)
	out = &model.PositionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.PositionInfo.Columns().Sort)
	// 执行查询
	// 指定item的键名用：ScanList
	if err := listModel.ScanList(&out.List, "Position"); err != nil {
		return out, err
	}
	// 不指定item的键名用：Scan
	//if err := listModel.Scan(&out.List); err != nil {
	//	return out, err
	//}
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
