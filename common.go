package e

type Err struct {
	Code int
	Msg  string
}

func (e *Err) Error() string {
	return e.Msg
}

var (
	COMMON_SUC        = &Err{10000, "success"}
	COMMON_BADREQUEST = &Err{10001, "请求错误，请更正后重试"}
	COMMON_PARAM_ERR  = &Err{10002, "请求参数错误，请更正后重试"}
	COMMON_PARAM_MISS = &Err{10003, "请求参数缺失，请补充后重试"}
	COMMON_NOT_FOUND  = &Err{10004, "未找到请求资源，请稍后重试"}
	COMMON_CONFILCT   = &Err{10005, "请求冲突，请稍后重试"}

	COMMON_INTERNAL_ERR             = &Err{15000, "内部错误，请稍后重试"}
	COMMON_INTERNAL_CALLING_ERR     = &Err{15001, "内部调用错误，请稍后重试"}
	COMMON_INTERNAL_CALLING_TIMEOUT = &Err{15002, "内部调用无响应，请稍后重试"}
	COMMON_INTERNAL_DB_ERR          = &Err{15003, "内部调用数据库错误，请稍后重试"}
	COMMON_INTERNAL_BADGATEWAY      = &Err{15004, "网关错误，请稍后重试"}
)

func NewBadRequestError(msg string) *Err {
	return &Err{COMMON_BADREQUEST.Code, msg}
}

func NewInternalError(msg string) *Err {
	return &Err{COMMON_INTERNAL_ERR.Code, msg}
}

func NewParamError(msg string) *Err {
	return &Err{COMMON_PARAM_ERR.Code, msg}
}

func NewParamMissError(msg string) *Err {
	return &Err{COMMON_PARAM_MISS.Code, msg}
}

func NewInternalDBError(msg string) *Err {
	return &Err{COMMON_INTERNAL_DB_ERR.Code, msg}
}

func NewInternalCallTimeoutError(msg string) *Err {
	return &Err{COMMON_INTERNAL_CALLING_TIMEOUT.Code, msg}
}
