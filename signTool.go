package tools

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/wenzhenxi/gorsa"
	"io"
	"strings"
)

// SignMD5withRSA md5-rsa加密 转base64
func SignMD5withRSA(data []byte, privateKey string) (string, error) {
	grsa := gorsa.RSASecurity{}
	grsa.SetPrivateKey(privateKey)
	md5Hash := md5.New()
	md5Hash.Write(data)
	hashed := md5Hash.Sum(nil)
	k, _ := grsa.GetPrivatekey()
	signByte, err := rsa.SignPKCS1v15(rand.Reader, k, crypto.MD5, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return sign, err
}

// SignMD5Up MD5加密转大写
func SignMD5Up(data string) string {
	w := md5.New()
	io.WriteString(w, data)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return strings.ToUpper(md5str2)
}
