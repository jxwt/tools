package tools

import (
	"testing"
)

func TestSingleSend(t *testing.T) {
	c := MessageSender{
		UserID: "账号",
		Pwd:    "密码",
	}
	ok := c.SendCheckCode("手机号", "333333", "30")
	t.Logf("%v", ok)
}
