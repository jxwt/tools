package tools

import (
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"path"
)

// FileWriteToPath 文件写入
func FileWriteToPath(filePath string, content string) bool {
	if !FileCreatePath(filePath) {
		return false
	}
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		logs.Error(err.Error())
		return false
	}
	return true
}

// FileCreatePath 创建文件夹
func FileCreatePath(filePath string, fileMode ...os.FileMode) bool {
	var mode os.FileMode = 0755
	if fileMode != nil && len(fileMode) > 0 {
		mode = fileMode[0]
	}
	err := os.MkdirAll(path.Dir(filePath), mode)
	if err != nil {
		logs.Error(err.Error())
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// FileImageToBase64 转换图片到BASE64字符串
func FileImageToBase64(file []byte) string {
	return base64.StdEncoding.EncodeToString(file) // 文件转base64
}

// FileExist 判断文件是否存在
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
