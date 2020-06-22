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

func SmsSend(smsURL string, mobile string, content string) *SmsSendResult {
	smsSendResult := new(SmsSendResult)
	result, _ := tools.HttpBeegoPost(smsURL, smsApiPostData(mobile, content), nil)
	json.Unmarshal(result, smsSendResult)

	return smsSendResult
}

func smsApiPostData(mobile string, content string) map[string]string {
	postData := make(map[string]string)
	postData["userid"] = "E101HS"
	postData["pwd"] = getSmsPassword("E101HS")
	postData["apikey"] = "596a36315dfc6e629099a427175acd6c"
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
