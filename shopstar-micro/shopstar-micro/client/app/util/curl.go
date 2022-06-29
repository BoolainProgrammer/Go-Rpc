package util

import (
	"io"
	"io/ioutil"
	"net/http"
)

/*
curl:
	请求：get / post / put / delete

	参数 - header
*/

type CurlInterface interface {
	Get() (string, error)
}

type Curl struct { // 它会进行修改或者调整，避免因为代码的改到导致项目的运行异常
	method string
	api    string
	body   io.Reader
	Header http.Header
}

func NewCurl(api string, body io.Reader) CurlInterface {
	return &Curl{
		api:  api,
		body: body,
	}
}

func (curl *Curl) Get() (string, error) {
	// 获取request请求
	req, err := http.NewRequest("GET", curl.api, curl.body)
	if err != nil {
		return "", err
	}
	// client就利用request去发送请求
	client := http.Client{}
	rsp,err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	// 处理信息
	body,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(body),nil

}
