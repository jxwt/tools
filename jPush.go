package tools

import (
	"fmt"
	"github.com/ylywyn/jpush-api-go-client"
)

const (
	appKey               = "0fbb9a2ab3d4b9a4ef4e3b87"
	secret               = "45603a23d94373b9fc71c904"
	inspectAppkey        = "7af735d1bddffe2e7211b9b1"
	inspectMastersecrete = "648be0a4820b16960d9c20dd"
)

// JPushSend 极光推送,audienceTags,在parking-cloud项目中取User.Id(即数据库表users的id字段)转字符串,一个可以发给多个user id;noticeAlert要推送的消息
func JPushSend(audienceTags []string, noticeAlert string) error {

	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)

	//Audience
	var ad jpushclient.Audience
	ad.SetTag(audienceTags)

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: noticeAlert})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: noticeAlert})

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)

	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}
	return err
}

func InspectJPushSendAlert(audienceTags []string, content string, extras map[string]interface{}) error {
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)

	var ad jpushclient.Audience
	ad.SetAlias(audienceTags)

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: content, Extras: extras})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: content, Extras: extras})

	//notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice from zy04031638"})
	//notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice from zy04031638"})
	//notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice "})
	//var msg jpushclient.Message
	//msg.Title = "Hello from zy0403"
	//msg.Content = "you are ylywn from zy0403"

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()

	//push
	c := jpushclient.NewPushClient(inspectMastersecrete, inspectAppkey)
	str, err := c.Send(bytes)

	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}
	return err
}

func InspectJPushSendMessage(audienceTags []string, noticeAlert string) error {

	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	//pf.Add(jpushclient.WINPHONE)
	//pf.All()

	//Audience
	var ad jpushclient.Audience
	//ad.SetTag(audienceTags)
	//s := []string{"19"}//s := []string{"1", "2", "3"}
	//ad.SetTag(s)
	ad.SetAlias(audienceTags)
	//ad.SetID(s)
	//ad.All()

	//Notice
	//var notice jpushclient.Notice
	//notice.SetAlert("alert")
	//notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: noticeAlert})
	//notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: noticeAlert})

	//notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice from zy04031638"})
	//notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice from zy04031638"})
	//notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice "})
	var msg jpushclient.Message
	//msg.Title = "Hello from zy0403"
	msg.Content = noticeAlert

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	//payload.SetNotice(&notice)
	payload.SetMessage(&msg)

	bytes, _ := payload.ToBytes()
	//push
	c := jpushclient.NewPushClient(inspectMastersecrete, inspectAppkey)
	str, err := c.Send(bytes)

	if err != nil {
		fmt.Printf("send %s err:%s", noticeAlert, err.Error())
	} else {
		fmt.Printf("send %s ok:%s", noticeAlert, str)
	}
	return err
}
