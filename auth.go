package e

var (
	AUTH_NOTLOGIN                 = &Err{40001, "user not login"}
	AUTH_NOPERMISSION             = &Err{40003, "user no permission"}
	AUTHINTERNAL_ROLE_NOT_DEFINED = &Err{40004, "role not defined"}
	AUTHINTERNAL_ROLE_NOT_SUPPORT = &Err{40005, "role not support"}
	AUTHINTERNAL_ROLE_INVALID     = &Err{40006, "role invalid"}
)
