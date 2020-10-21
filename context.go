package tools

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func ContextGetToken(auth *AuthTool, ctx *context.Context) string {
	token := ""
	if ctx.Request.Header.Get("Authorization") != "" {
		token = ctx.Request.Header.Get("Authorization")
	} else if ctx.Request.Header.Get("Token") != "" {
		token = ctx.Request.Header.Get("Token")
	} else if ctx.Request.Header.Get("token") != "" {
		token = ctx.Request.Header.Get("token")
	} else {
		cookie, _ := ctx.GetSecureCookie(string(auth.Key), "Authorization")
		token = cookie
	}

	return token
}

func CtxToJson(ctx *context.Context, obj interface{}) error {
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	if len(body) != 0 {
		err := json.Unmarshal(body, obj)
		return err
	}

	var values url.Values
	if ctx.Request.Method == "GET" {
		values  = ctx.Request.URL.Query()
	} else if ctx.Request.Method == "POST" {
		values = ctx.Request.PostForm
	}
	t := reflect.TypeOf(obj).Elem()
	params := make(map[string]interface{})
	if len(values) == 0 {
		for k, v := range ctx.Input.Params() {
			params[k] = parseObjectValue(t, k, v)
		}
	} else {
		for k, v := range values {
			params[k] = parseObjectValue(t, k, v[0])
		}
	}
	jsonData, err := json.Marshal(params)
	err = json.Unmarshal(jsonData, obj)

	return err
}

func parseObjectValue(t reflect.Type, k string, v string) interface{} {
	var val interface{} = v
	if strings.ToUpper(k) == "ID" {
		val, _ = strconv.Atoi(v)
		return val
	}
	for index := 0; index < t.NumField(); index++ {
		if strings.ToUpper(t.Field(index).Name) == strings.ToUpper(k) || t.Field(index).Tag.Get("json") == k {
			kind := t.Field(index).Type.Kind()
			if kind == reflect.Float64 {
				val = StringToFloat64(v)
			} else if t.Field(index).Type == reflect.TypeOf(&JSONTime{}) {
				val = v
			} else if kind != reflect.String {
				val, _ = strconv.Atoi(v)
			}
			break
		}
	}

	return val
}
