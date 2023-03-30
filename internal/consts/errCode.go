package consts

// OK 成功返回
const OK int = 200

// 全局错误码 (前3位代表业务, 后三位代表具体功能)
const (
	ServerCommonError         int = 100001
	RequestParamError         int = 100002
	TokenExpireError          int = 100003
	TokenGenerateError        int = 100004
	DbError                   int = 100005
	DbUpdateAffectedZeroError int = 100006
	ParamValidateError        int = 100007
	PermissionDenied          int = 100008
	JsonMarshalError          int = 100009
	AsynqEnqueueError         int = 100010
	ElasticSearchError        int = 100011
	RequestLimitError         int = 100012
)

// 业务相关错误码
const (
	AdminNotFound  int = 200001
	AdminNameExist int = 200002
)
