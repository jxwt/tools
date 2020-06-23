package tools

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"time"
)

type SmsSendResult struct {
	Result int    `json:"result"`
	Msgid  int64  `json:"msgid"`
	Custid string `json:"custid"`
}

type SmsSender struct {
	UserID string
	ApiKey string
}

func (s *SmsSender) SmsSend(smsURL string, mobile string, content string) *SmsSendResult {
	smsSendResult := new(SmsSendResult)
	result, _ := HttpBeegoPost(smsURL, s.smsApiPostData(mobile, content), nil)
	json.Unmarshal(result, smsSendResult)

	return smsSendResult
}

func (s *SmsSender) smsApiPostData(mobile string, content string) map[string]string {
	postData := make(map[string]string)
	postData["userid"] = s.UserID
	postData["pwd"] = getSmsPassword(s.UserID)
	postData["apikey"] = s.ApiKey
	postData["mobile"] = "0086" + mobile
	postData["content"] = GbkToUtf8(content)
	//postData["content"]= content
	logs.Warning(postData)
	return postData
}

func getSmsPassword(userId string) string {
	encodeString := userId + "000000006OsCcC" + GetTimeNumberByFormat(time.Now(), "mdHis")
	return Md5(encodeString)
}
