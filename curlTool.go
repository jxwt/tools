package tools

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"strings"
)

// GetRequestAgentType 从报文获取是手机或是pc
func GetRequestAgentType(ctx *context.Context) string {
	keywords := []string{"Android", "iPhone", "iPod", "iPad", "Windows Phone", "MQQBrowser"}
	for _, keyword := range keywords {
		if strings.Contains(ctx.Request.UserAgent(), keyword) {
			return "phone"
		}
	}
	return "pc"
}

// HTTPBeegoJSONPost http json post
func HTTPBeegoJSONPost(url string, jsonBody string, header map[string]string) ([]byte, error) {
	req := httplib.Post(url)
	req.Body(jsonBody)
	for k, v := range header {
		req.Header(k, v)
	}
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// HTTPBeegoQueryPost http query post
func HTTPBeegoQueryPost(url string, params map[string]string, header map[string]string) string {
	req := httplib.Post(url)

	for k, v := range params {
		req.Param(k, v)
	}
	for k, v := range header {
		req.Header(k, v)
	}
	resp, err := req.Response()
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// HTTPBeegoGet http get
func HTTPBeegoGet(url string, header map[string]string) ([]byte, error) {
	req := httplib.Get(url)
	for k, v := range header {
		req.Header(k, v)
	}
	resp, err := req.Response()
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		logs.Warn(err.Error())
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
