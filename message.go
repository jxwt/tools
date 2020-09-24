package tools

import (
	"fmt"
	"github.com/jxwt/tools/montnets/mwgate/smsutil"
)

type MessageSender struct {
	UserID string // 用户
	Pwd    string // 密码或者授权码
}

// CodeTemp 验证码模板
const CodeTemp = "您的验证码是%s，在%s分钟内输入有效。如非本人操作请忽略此短信。"

func singleSend(userid string, pwd string, mobile string, content string) bool {
	// 将数据打包
	sendobj := smsutil.NewSingleSend(userid, pwd, mobile, content)
	// 发送数据
	return smsutil.SendAndRecvOnce(sendobj)
}

// SendCheckCode 发送验证码
// period 有效期
func (c *MessageSender) SendCheckCode(mobile, code, period string) bool {
	content := fmt.Sprintf(CodeTemp, code, period)
	return singleSend(c.UserID, c.Pwd, mobile, content)
}
