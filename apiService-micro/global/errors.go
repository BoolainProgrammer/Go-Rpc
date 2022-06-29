package global

import "errors"

var (
	ErrGormRecordNotFound = errors.New("查询不到信息")
	ErrGormRecord         = errors.New("查询异常")
)

var (
	ErrApiNameExist = errors.New("服务名称已存在")
	ErrApiUrlExist  = errors.New("服务链接已存在")
	ErrCreateApi    = errors.New("服务链接创建失败")
	ErrUpdateApi    = errors.New("服务链接更新失败")
)
var (
	ErrMenuNameExist = errors.New("菜单名称已存在")
	ErrMenuCodeExist = errors.New("菜单代码已存在")
	ErrUpdateMenu = errors.New("菜单更新失败")
)


var (
	ErrJwtAuthFailure = errors.New("您的令牌失效")
	ErrJwtAuthIllegal = errors.New("未登录或非法访问")
	ErrJwtAuthExpired = errors.New("令牌过期")
)
