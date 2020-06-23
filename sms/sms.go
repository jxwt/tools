package sms

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/jxwt/tools"
	"time"
)

type SmsSendResult struct {
	Result int    `json:"result"`
	Msgid  int64  `json:"msgid"`
	Custid string `json:"custid"`
}

type Sender struct {
	UserID string
	ApiKey string
}

func (s *Sender) SmsSend(smsURL string, mobile string, content string) *SmsSendResult {
	smsSendResult := new(SmsSendResult)
	result, _ := tools.HttpBeegoPost(smsURL, s.smsApiPostData(mobile, content), nil)
	json.Unmarshal(result, smsSendResult)

	return smsSendResult
}

func (s *Sender) smsApiPostData(mobile string, content string) map[string]string {
	postData := make(map[string]string)
	postData["userid"] = s.UserID
	postData["pwd"] = getSmsPassword(s.UserID)
	postData["apikey"] = s.ApiKey
	postData["mobile"] = "0086" + mobile
	postData["content"] = tools.GbkToUtf8(content)
	//postData["content"]= content
	logs.Warning(postData)
	return postData
}

func getSmsPassword(userId string) string {
	encodeString := userId + "000000006OsCcC" + tools.GetTimeNumberByFormat(time.Now(), "mdHis")
	return tools.Md5(encodeString)
}
