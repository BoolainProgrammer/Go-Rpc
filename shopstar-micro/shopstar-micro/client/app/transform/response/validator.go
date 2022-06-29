package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"sixstar/shopstar-micro/client/core/validate"
	"strings"
)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

func HandleValidatorError(ctx *gin.Context, err error) {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)

	if !ok {
		FailWithDetailed(errs, "校验失败", ctx)
		return
	}

	FailWithDetailed(removeTopStruct(errs.Translate(validate.Trans)), "校验失败", ctx)
}
// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		// 从文本的逗号开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		if err == "" {
			err = validate.Err.String()
		}

		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
