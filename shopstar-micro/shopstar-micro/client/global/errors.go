package global

import "errors"

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */



var (
	ErrUsernameExist        = errors.New("用户名以存在")

	ErrGormRecordNotFound   = errors.New("查询不到信息")
	ErrGormRecord           = errors.New("查询异常")

	ErrPassword             = errors.New("密码错误")
	ErrPasswordHash         = errors.New("加密错误")
	ErrRegister				= errors.New("注册失败,请重试")

	ErrSkuNum				= errors.New("该商品库存不足")
)

var (
	ErrValidatorSmsCode    = errors.New("短信验证错误,请稍后再试")
	ErrValidatorCaptahCode = errors.New("验证码错误,请稍后再试")
)

var (
	ErrJwtAuthFailure		= errors.New("您的令牌失效")
	ErrJwtAuthIllegal		= errors.New("未登录或非法访问")
	ErrJwtAuthExpired		= errors.New("令牌过期")
)

