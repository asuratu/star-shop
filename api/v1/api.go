package v1

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，默认10，最大50"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数"`
	Page  int `dc:"分页码"`
	Size  int `dc:"分页数量"`
}
