package tools

import (
	"encoding/json"
	"errors"
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

// HttpBeegoJsonPost http json post
func HttpBeegoJsonPost(url string, jsonBody string, header map[string]string) ([]byte, error) {
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

// HttpBeegoPost http query post
func HttpBeegoPost(url string, params map[string]string, header map[string]string) ([]byte, error) {
	req := httplib.Post(url)

	for k, v := range params {
		req.Param(k, v)
	}
	for k, v := range header {
		req.Header(k, v)
	}
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// HttpBeegoGet http get
func HttpBeegoGet(url string, header map[string]string) ([]byte, error) {
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


func HttpRequest2Map(ctx *context.Context) (map[string]interface{}, error) {
	if ctx.Request.Method == "GET" {
		return httpForm2Map(ctx)

	} else if ctx.Request.Method == "POST" {
		return httpBody2Map(ctx)

	}

	return nil, errors.New("unSupport http method")
}

func httpBody2Map(ctx *context.Context) (map[string]interface{}, error) {
	r := ctx.Request
	m := make(map[string]interface{})
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func httpForm2Map(ctx *context.Context) (data map[string]interface{}, err error) {
	r := ctx.Request
	data = make(map[string]interface{})
	err = r.ParseForm()
	if err != nil {
		return
	}

	if len(r.Form) == 0 {
		return nil, errors.New("输入参数为空")
	}
	for i, v := range r.Form {
		if len(v) == 0 || len(v) > 1 {
			return nil, errors.New("表单输入不合法")
		}
		data[i] = v[0]
	}
	return
}
