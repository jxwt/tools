package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func DataToStruct(data map[string]string, out interface{}) error {
	ss := reflect.ValueOf(out).Elem()
	for i := 0; i < ss.NumField(); i++ {
		val := data[ss.Type().Field(i).Name]
		name := ss.Type().Field(i).Name
		switch ss.Field(i).Kind() {
		case reflect.String:
			ss.FieldByName(name).SetString(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetInt(int64(i))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetUint(uint64(i))
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetFloat(f)
		default:
		}
	}
	return nil
}

// Map2Struct map转struct
func Map2Struct(data map[string]string, out interface{}) error {
	ss := reflect.ValueOf(out).Elem()
	for i := 0; i < ss.NumField(); i++ {
		val := data[ss.Type().Field(i).Name]
		name := ss.Type().Field(i).Name
		switch ss.Field(i).Kind() {
		case reflect.String:
			ss.FieldByName(name).SetString(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetInt(int64(i))
		case reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetUint(uint64(i))
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				continue
			}
			ss.FieldByName(name).SetFloat(f)
		default:
		}
	}
	return nil
}

// Struct2Map struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}

	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}

	}
	return data
}

// Struct2Json2Map .
func Struct2Json2Map(obj interface{}) map[string]interface{} {
	bytes, _ := json.Marshal(obj)
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(bytes), &mapResult)
	return mapResult
}

// Interface2String .
func Interface2String(value interface{}) string {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.String:
		return value.(string)
	case reflect.Int:
		return strconv.Itoa(value.(int))
	case reflect.Int64:
		return strconv.FormatInt(value.(int64), 10)
	case reflect.Uint32:
		v := uint64(value.(uint32))
		return strconv.FormatUint(v, 10)
	case reflect.Uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case reflect.Float32:
		v := float64(value.(float32))
		return strconv.FormatFloat(v, 'f', 2, 64)
	case reflect.Float64:
		return strconv.FormatFloat(value.(float64), 'f', 2, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.(bool))
	default:
		return ""
	}
}

// MapToSign .
func MapToSign(param map[string]interface{}) string {
	data := make(map[string]string)
	var keys []string
	for k, v := range param {
		if k != "" && v != nil && v != "" {
			keys = append(keys, k)
			switch v.(type) {
			case string:
				data[k] = fmt.Sprintf(`%s=%s`, k, v)
				break
			case int:
				data[k] = fmt.Sprintf(`%s=%d`, k, v)
				break
			case int64:
				data[k] = fmt.Sprintf(`%s=%d`, k, v)
				break
			case uint:
				data[k] = fmt.Sprintf(`%s=%d`, k, v)
			case float64:
				float := strconv.FormatFloat(v.(float64), 'f', -1, 64)
				data[k] = fmt.Sprintf(`%s=%s`, k, float)
				break
			case bool:
				data[k] = fmt.Sprintf(`%s=%t`, k, v)
				break
			default:
				marshal, _ := json.Marshal(v)
				data[k] = fmt.Sprintf(`%s=%s`, k, string(marshal))
			}
		}
	}
	sort.Strings(keys)
	var sign []string
	for _, key := range keys {
		sign = append(sign, data[key])
	}
	signData := strings.Join(sign, "&")
	return signData
}

func Struct2JsonMap(obj interface{}) map[string]interface{} {
	bytes, _ := json.Marshal(obj)
	var mapResult map[string]interface{}
	_ = json.Unmarshal(bytes, &mapResult)
	return mapResult
}
