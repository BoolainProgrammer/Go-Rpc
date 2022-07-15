package sms

import (
	"fmt"
	"net/url"
	"sixstar/shopstar-micro/client/app/util"
	"strings"

)



type Smsbao struct{}

func (sms *Smsbao) Send(phone, content string) (string, error) {
	// 1. 拼接url
	//var sendapi strings.Builder
	//sendapi.Grow(200)
	//sendapi.WriteString()
	//sendapi.WriteString("?u=")
	//sendapi.WriteString(global.Config.Smsbao.User)
	//sendapi.WriteString("&p=")
	//sendapi.WriteString(util.Md5([]byte(global.Config.Smsbao.Pass)))
	//sendapi.WriteString("&m=")
	//sendapi.WriteString(phone)
	//sendapi.WriteString("&c=")
	//sendapi.WriteString(url.QueryEscape(content))
	//// 2. get请求
	//code, err := util.NewCurl(sendapi.String(), nil).Get()
	//// 3. 处理返回信息
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}
	//
	//return global.Config.Smsbao.StatusStr[code], nil
}
