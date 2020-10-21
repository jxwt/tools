package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"net/url"
	"strings"
	"time"
)

const (
	singleUrl = "http://api01.monyun.cn:7901/sms/v2/std/single_send"
	batchUrl = "http://api01.monyun.cn:7901/sms/v2/std/batch_send"
)
type MessageSender struct {
	Userid string `json:"userid"`
	Pwd string `json:"pwd"`
	Timestamp string `json:"timestamp"`
	Mobile string `json:"mobile"`
	Content string `json:"content"`
}

type Result struct {
	Result int `json:"result"`
	MsgId  int64 `json:"msgid"`
	CustId string 	`json:"custid"`
}

// CodeTemp 验证码模板
const CodeTemp = "您的验证码是%v，在%v分钟内输入有效。如非本人操作请忽略此短信。"


func FormatContent(content string) string {
	// 去掉两端的空格
	content = strings.TrimSpace(content)
	// 转为GBK
	s := mahonia.NewEncoder("gbk").ConvertString(content)
	v := url.Values{}
	v.Set("aa", s)
	str := v.Encode()
	arr := strings.Split(str, "=")
	return arr[1]
}

func (i *MessageSender) send(url string) error {
	md5Ctx := md5.New()
	i.Timestamp = time.Now().Format("060102150405")[2:]
	userid := strings.ToUpper(i.Userid)
	md5Ctx.Write([]byte(userid + "00000000" + i.Pwd + i.Timestamp))
	encryptPwd := md5Ctx.Sum(nil)
	i.Pwd = hex.EncodeToString(encryptPwd[:])
	i.Content = FormatContent(i.Content)
	d, _ := json.Marshal(i)
	header := make(map[string]string)
	header["Accept-Encoding"] = "gzip, deflate"
	header["Content-Type"] = "application/json"
	header["Connection"] =  "Close"
	ret, err := HttpBeegoJsonPost(url, string(d), header)
	if err != nil {
		panic(err)
	}
	r := new(Result)
	err = json.Unmarshal(ret, r)
	if err != nil {
		return err
	}
	if r.Result != 0 {
		return errors.New("发送失败")
	}
	return nil
}

//// SendCheckCode 发送验证码
//// content 短信内容 可不传
//// args 0 是模版 1 开始时数值
func (i *MessageSender) Send(mobile []string, temp string, args ...interface{}) error {
	if len(mobile) == 0 {
		return errors.New("手机号不能为空")
	}
	if temp == "" && len(args) == 0 {
		return errors.New("短信内容不能为空")
	}
	if temp == "" {
		temp = CodeTemp
	}
	i.Content = fmt.Sprintf(temp, args...)
	if len(mobile) == 1 {
		i.Mobile = mobile[0]
		return i.send(singleUrl)
	} else {
		i.Mobile = strings.Join(mobile, ",")
		return i.send(batchUrl)
	}

	return nil
}


