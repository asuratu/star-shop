package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Api 定义业务侧的数据结构, 提供对外接口的输入|输出参数定义

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
