package request
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type SmsRequest struct {
	Phone string `form:"phone" binding:"required,phone"`
}
