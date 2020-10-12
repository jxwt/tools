package tools

import (
	"github.com/astaxie/beego/logs"
	"github.com/ylywyn/jpush-api-go-client"
)

// JPushSender .
type JPushSender struct {
	AppKey    string
	SecretKey string
}

// JPushSend 推送消息
// audienceTags 推送目标 使用tag标签(用户ID数组)
// alert 通知
// 返回 发送内容,错误
func (s *JPushSender) JPushSend(audienceTags []string, alert string) (string, error) {
	var pf jpushclient.Platform
	pf.All()

	// Audience 推送目标
	var ad jpushclient.Audience
	ad.SetTag(audienceTags)

	// Notice 通知
	var notice jpushclient.Notice
	notice.SetAlert(alert)
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)      // 设置平台
	payload.SetAudience(&ad)      // 设置目标
	payload.SetNotice(&notice)    // 设置通知
	bytes, _ := payload.ToBytes() // json化

	//push
	c := jpushclient.NewPushClient(s.SecretKey, s.AppKey)
	str, err := c.Send(bytes)
	if err != nil {
		logs.Warning("JPushSend err", err.Error())
		return str, err
	}
	return str, nil
}
