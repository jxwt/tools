package tools

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthTool struct {
	Key            []byte
	ExpiresSeconds int64
	Claims         jwt.MapClaims
}

func NewAuthTool() *AuthTool {
	return new(AuthTool)
}

func (i *AuthTool) SetKey(keys ...string) *AuthTool {
	if len(keys) == 0 {
		i.Key = []byte("jhsno1")
	} else {
		tmp := ""
		for _, key := range keys {
			tmp += key
		}
		i.Key = []byte(tmp)
	}
	return i
}

func (i *AuthTool) SetExpiresSeconds(seconds int64) *AuthTool {
	if seconds == 0 {
		i.ExpiresSeconds = 1000
	} else {
		i.ExpiresSeconds = seconds
	}
	return i
}

func (i *AuthTool) MakeClaims(data map[string]interface{}) *AuthTool {
	i.Claims = make(jwt.MapClaims)
	i.Claims["exp"] = int64(time.Now().Unix() + i.ExpiresSeconds)
	for k, v := range data {
		i.Claims[k] = v
	}
	return i
}

func (i *AuthTool) GenToken(data map[string]interface{}) string {
	i.MakeClaims(data)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, i.Claims)
	ss, err := token.SignedString(i.Key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}

// CheckToken 校验token是否有效
func (i *AuthTool) CheckToken(token string) error {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return i.Key, nil
	})
	return err
}

// ParseKey 解析token
func (i *AuthTool) ParseKey(token string) *jwt.Token {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return i.Key, nil
	})
	if err != nil {
		return nil
	}
	return t
}
