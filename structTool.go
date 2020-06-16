package tools

import (
	"fmt"
	"reflect"
	"strconv"
)

// StructFloatFormat 格式化结构体内的float 保留2位小数
func StructFloatFormat(obj interface{}) interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		t = t.Elem()
		v = v.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Interface().(type) {
		case float64:
			formatFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.Field(i).Interface().(float64)), 64)
			v.Field(i).Set(reflect.ValueOf(formatFloat))
		default:
			continue
		}
	}
	return obj
}
