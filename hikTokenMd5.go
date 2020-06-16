package tools

import (
	"encoding/json"
	"strings"
)

// CreateHikReqToken .
func CreateHikReqToken(uri string, param map[string]interface{}, secret string) string {
	paramStr, _ := json.Marshal(param)
	return strings.ToUpper(Md5(uri + string(paramStr) + secret))
}
