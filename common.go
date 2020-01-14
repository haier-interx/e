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
	COMMON_BADREQUEST = &Err{10001, "bad request"}
	COMMON_PARAM_ERR  = &Err{10002, "request parameter error"}
	COMMON_PARAM_MISS = &Err{10003, "request parameter missing"}
	COMMON_NOT_FOUND  = &Err{10004, "not found"}
	COMMON_CONFILCT   = &Err{10005, "conflict"}

	COMMON_INTERNAL_ERR             = &Err{15000, "unknown internal error"}
	COMMON_INTERNAL_CALLING_ERR     = &Err{15001, "internal calling error"}
	COMMON_INTERNAL_CALLING_TIMEOUT = &Err{15002, "internal calling timeout"}
	COMMON_INTERNAL_DB_ERR          = &Err{15003, "internal database error"}
	COMMON_INTERNAL_BADGATEWAY      = &Err{15004, "bad gateway"}
)
