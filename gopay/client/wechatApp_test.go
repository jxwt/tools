package client

import (
	"parking_mall/utils/gopay/common"
	"testing"
)

// TestMicroPay 微信统一收单测试
func TestMicroPay(t *testing.T) {
	wxPayTool := WechatAppClient{
		AppID: "wxeec1af7e2251f967",
		MchID: "1510147771",
		Key:   "0tOP61gTCQRZWKt62HGY6cNButHL04AL",
	}
	request := common.MicroPayRequest{
		OutTradeNo: "WxPay20f9062010dsd1J0N1jinZL7",
		TotalFee:   1,
		AuthCode:   "135014210759658562",
	}
	res, err := wxPayTool.MicroPay(&request)
	if err != nil {
		t.Errorf("err:%v\n", err)
	}
	// 需要输入密码
	if res.ReturnCode == "SUCCESS" && res.ErrCode == "USERPAYING" {
		t.Errorf("11111111111\n")
		//TODO 扫单
	}
	t.Errorf("response2:%v\n", res)
}

// TestQueryOrder 查询订单测试
func TestQueryOrder(t *testing.T) {
	wxPayTool := WechatAppClient{
		AppID: "wxeec1af7e2251f967",
		MchID: "1510147771",
		Key:   "0tOP61gTCQRZWKt62HGY6cNButHL04AL",
	}
	res, err := wxPayTool.QueryOrder("WxPay20f9062010dsd1J0N1jinZL7")
	if err != nil {
		t.Errorf("err:%v\n", err)
	}
	// 扫单情况 1
	if res.ReturnCode == "SUCCESS" && res.ErrCode == "USERPAYING" {
		t.Errorf("11111111111\n")
		//TODO 扫单
	}
	// 扫单情况 2
	if res.ReturnCode == "SUCCESS" && res.ErrCode == "PAYERROR" {
		t.Errorf("22222222222\n")
		//TODO 撤单 撤单后返回状态 REVOKED
	}
	t.Errorf("response2:%v\n", res)
}
