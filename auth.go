package e

var (
	AUTH_NOTLOGIN                 = &Err{40001, "您未登录或登录过期，请登录"}
	AUTH_NOPERMISSION             = &Err{40003, "您无权访问"}
	AUTHINTERNAL_ROLE_NOT_DEFINED = &Err{40004, "用户角色不可用，请联系系统管理员更正"}
	AUTHINTERNAL_ROLE_NOT_SUPPORT = &Err{40005, "无此角色信息，请检查配置权限是否正确"}
	AUTHINTERNAL_ROLE_INVALID     = &Err{40006, "用户角色不可用"}
)
