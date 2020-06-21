package tools

import (
	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func EncryptPassword(password string) string {
	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		logs.Error(err)
	}
	return string(cryptPassword)
}

// 获取Token
func GetAuthCode() int {
	return GetRandIntn(99999, 10000)
}

// 比较密码
func MatchPassword(dst string, src string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dst), []byte(src)) == nil
}

