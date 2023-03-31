package xerr

import (
	"shop/internal/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var message map[int]string

func init() {
	message = make(map[int]string)
	// 全局错误码
	message[consts.OK] = "SUCCESS"
	message[consts.ServerCommonError] = "服务器开小差啦，稍后再来试一试"
	message[consts.RequestParamError] = "参数错误"
	message[consts.TokenExpireError] = "token失效，请重新登陆"
	message[consts.TokenGenerateError] = "生成token失败"
	message[consts.DbError] = "数据库繁忙，请稍后再试"
	message[consts.DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[consts.ParamValidateError] = "参数校验失败"
	message[consts.PermissionDenied] = "权限不足"
	message[consts.JsonMarshalError] = "json marshal fail"
	message[consts.AsynqEnqueueError] = "异步任务入队失败"
	message[consts.ElasticSearchError] = "es操作失败"
	message[consts.RequestLimitError] = "请求过于频繁，请稍后再试"

	// 用户模块
	message[consts.AdminNotFound] = "用户不存在"
	message[consts.AdminNameExist] = "用户已存在"
	message[consts.AdminNameOrPasswordError] = "用户名或密码错误"
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode int) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}

func WrapErr(errcode int) error {
	msg := MapErrMsg(errcode)
	return gerror.NewCode(gcode.New(errcode, msg, nil), msg)
}
