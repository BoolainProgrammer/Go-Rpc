package v1

import (
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/gin-gonic/gin"

	"sixstar/shopstar-micro/client/global"
	"sixstar/shopstar-micro/client/core/cache"
	"sixstar/shopstar-micro/client/app/util"
	"sixstar/shopstar-micro/client/app/transform/request"
	"sixstar/shopstar-micro/client/app/transform/response"

)

type Verification struct {
	
}

func (*Verification) SmsCodes(ctx *gin.Context) {
	// 1. 获取手机号码 phone
	smsreq := new(request.SmsRequest)
	// 2. phone 信息是需要进行验证判断是否为正常的手机号码
	if err := ctx.ShouldBind(smsreq); err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}

	// 3. 生成发送信息
	code := util.GenValidateCode(global.Config.Sms.Long)
	content := fmt.Sprintf(global.Config.Sms.Temp.Code, code, global.Config.Sms.Overdue)
	// 4. 发送短信
	fmt.Println(content)
	fmt.Println("code : ", code)
	//_, err := sms.NewSms(global.Config.Sms.Service).Send(smsreq.Phone, content)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), ctx)
	//	return
	//}

	// 5. 存储
	key := util.GenValidateCode(15) // 多种方案
	cache.CacheManager.Set(key, []byte(code), &cache.Options{
		Expiration: time.Duration(global.Config.Sms.Overdue) * time.Second,
	})

	// 6. 通知
	response.OkWithData(&gin.H{
		"verification_key": key,
		"expired_ad":       global.Config.Sms.Overdue,
	}, ctx)
}

func (*Verification) Captcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.Config.Captche.ImgHeight, global.Config.Captche.ImgWidth, global.Config.Captche.KeyLong, 0.7, 80)

	//driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	cp := base64Captcha.NewCaptcha(driver, global.CaptchaStore)

	id, b64s, err := cp.Generate()

	if err != nil {
		//global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMsg("验证码获取失败", ctx)
		return
	}

	response.OkWithDetailed(response.CaptchaResp{
		CaptchaId: id,
		PicPath:   b64s,
	}, "验证码获取成功", ctx)
}
