package response

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	ERRORCREATE = "创建失败"
	ERROROPTION = "操作失败"
	ERRORGIN	= "系统异常"

	SUCCESS = 200
	SUCCESSCREATE = "创建成功"
	SUCCESSOPTION = "操作成功"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMsg(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMsg(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
func FailWithErr(err error, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, err.Error(), c)
}


func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}