package client

import (
	"crypto/rsa"
	"github.com/axgle/mahonia"
	"parking_mall/utils/gopay/common"
	"testing"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// TestToaccountTransferRequest 转账测试
func TestToaccountTransferRequest(t *testing.T) {
	// 私钥
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtauMp34N+g+hyrD6845SBaatoU7iSzH/pcyRwaW07z/kREYtpCkpPOTo4O570TR5pexQyhMbg+oi87EnO2dUJkOzkj6rAMcioNZMeXY27AjmbxHy+LwvPSl3tF6nZDU1s+DGEJzM2imdNoa2h117STUs9kwE7c+SE6E3JuDrmqI+4+DkPOJBmgwnN7TVU5AC+Cckv2Uv1shzG1P1P/1/2AzLKqh7ghjgy6uQgsHvU5zaGawS3KtFmVXGiBCekbMrhvo98LRv40D9mYOBsW8t9kLBQMIleNiU4aP+6JzGDRbrnGM0vgYfQKe9C7IK9X/JQEwlpIOI6d6qraCONW1brAgMBAAECggEAEs2eJ/ImTdd7osNuYgjDF20fusYpIzGtRODJOK8VuwCH3wPp+8+z0vc/is1cJN0fyQwhWoDvF4HSxblRH26bHNhr9zRkrUY4nZSBiS9XqMlK+crKQ8zSGP3VRVnlfrVkicY3iD6/3NAQ92fqbbvuehsLZ3fDEHE2e/q8RH0HVe8O9Ejh1wVTa9MMjd4+nSIt06U+6cDfdPZGLWynPqrMaC4Jtk8U/TRA2EGDK84vK5mrqbim737iXNB8usQ+lGXDlsHVrVDG/9mP8cFW6AoqVRxwN8FixG+xGZFUcnIxAm5kSY6nE86TS+61v6oajTGLbNakc2tozROjRD6uZqzO4QKBgQDs+rNgKD3+fcamN48hXqC7CwQvg1WxyuV0VkK4MKXMbJDwmbxLSku+xh0P9920CjAupcPcsKFAj2FPkEdbkA9Wmm5wTlh9XqauDuKVAkqd77DwkvB4WoY2bQGov16qkQ420uvDc+DysbE7cUvVcl8qcn7YcSO0NsM9+PYlb1GxuQKBgQC7Vij1cDIIrltkx3pggNyiy+nEpx4uhxuDfO8xVSvI5FwSDExPyza50isDKWxJow5iUw9Rz2ybJ2emBeLKwm4G0Cj7Zl/001nsAevu/ob4Hip9qSl+0+Z+Idmvd6BFEYJHoRNK7FjHdwXdBNydwU8BtZZzU3/yHzZ3OIkZFoAvwwKBgQCY6H67advujNuTvr+1CWjup4IQ4k52BPQfJ9WvIXyptdej156+efb75rsz5XyBQh2qy7zgdnvlu15Px0mz0/WBrO3bu0Gvy1YDc4lSGoNo+xMRd85/6fE1xwpOBwUfS69/QoNrvyaDkpJIR6dl14F+UxhzsjUWgEtkfnLc3sI4yQKBgBUMChABKe3lwOjirHIZKDC2Hi505CQwE7xDFhCB1Ch+14VDknNIjn50CVcSmVLwmdYcJNV5K2eHFtMFSESlcX0cd+4+wzsbX7fvQ1WXjQxlPzrc/Yd9QSEcpntbQktgOzXW9/br9NF8ItGBEVQ7+qdjgmK0l+RY82KTnHuQFpjHAoGBAN450ijHPUn67IP9r/KyGrtr21y7Oca1EDpKXbM6p3/H7/637ygf08iJpcYdsh7cslkIVOBxmvPYJZTHlCCD2HoSbEMdwou18rrTst0P72772fayLq6JDvjSjR+fK0blU3RXBlvPAnpi5K/YducknZWitAynC6XZEIytKLDZFTsF"
	bytePrivateKey := FormatPrivateKey(privateKey)
	pri, err := ParsePKCS1PrivateKey(bytePrivateKey)
	if err != nil {
		t.Errorf("ParsePKCS1PrivateKey err:%v\n", err)
		return
	}
	// 公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAma31S+2cqIDZ9jb3kasYR3kVkzCs0r2hw53WX+AqNOWdIRyT762diusA6Xn/EExJ1ph1C9rWBUQy+GaM5blqz1QZbsXfVN3ITJFtWxhPN+jkKqZjdQ1mrGRxwTlcCGFbRy2FhdxsL4x2tsZDVh4P05/+3ZYuVQv2ep8O69DfEVMFbSmXUif73c2eSOthXV8OLN+gar2yhuIAkZoui35cdklfmtV3gNzNO+ykZrIVTzJr4/kIQVlxBRexdFZWgXrE35Kji8ql9p/lF/1v0tXsH0KHqg9L4HIGZ69yEXojdnplZ0h9Is4TRpFlBCDUZFS4AXJqLKQKd+lg1DkEuh6gowIDAQAB"
	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = ParsePKCS1PublicKey(FormatPublicKey(aliPublicKey))
		if err != nil {
			t.Errorf("ParsePKCS1PublicKey err:%v\n", err)
			return
		}
	}
	aliAppClient := new(AliAppClient)
	aliAppClient.PublicKey = pub
	aliAppClient.PrivateKey = pri
	aliAppClient.AppID = "2018030602325750"

	payRefundRequest := common.ToaccountTransferRequest{
		OutBizNo:      "njkdvji3o3e93dfdsjihf9j21ggf3",
		PayeeType:     "ALIPAY_LOGONID",
		PayeeAccount:  "15070590208",
		Amount:        "0.1",
		PayeeRealName: "黄晨帆",
		Remark:        "测试2",
		PayerShowName: "逐一科技",
	}
	res, err := aliAppClient.MakeToaccountTransfer("alipay.fund.trans.toaccount.transfer", &payRefundRequest, "RSA2")
	if err != nil {
		t.Errorf("MakeToaccountTransfer err:%v\n", err)
	}
	response, err := aliAppClient.SendToAlipay(res, "post")
	if err != nil || response == "" {
		t.Errorf("SendToAlipay err:%v\n", err)
	}
	response = ConvertToString(response, "gbk", "utf-8")
	t.Errorf("response:%v\n", response)
}

// TestRefund 支付宝退款测试
func TestRefund(t *testing.T) {
	// 私钥
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtauMp34N+g+hyrD6845SBaatoU7iSzH/pcyRwaW07z/kREYtpCkpPOTo4O570TR5pexQyhMbg+oi87EnO2dUJkOzkj6rAMcioNZMeXY27AjmbxHy+LwvPSl3tF6nZDU1s+DGEJzM2imdNoa2h117STUs9kwE7c+SE6E3JuDrmqI+4+DkPOJBmgwnN7TVU5AC+Cckv2Uv1shzG1P1P/1/2AzLKqh7ghjgy6uQgsHvU5zaGawS3KtFmVXGiBCekbMrhvo98LRv40D9mYOBsW8t9kLBQMIleNiU4aP+6JzGDRbrnGM0vgYfQKe9C7IK9X/JQEwlpIOI6d6qraCONW1brAgMBAAECggEAEs2eJ/ImTdd7osNuYgjDF20fusYpIzGtRODJOK8VuwCH3wPp+8+z0vc/is1cJN0fyQwhWoDvF4HSxblRH26bHNhr9zRkrUY4nZSBiS9XqMlK+crKQ8zSGP3VRVnlfrVkicY3iD6/3NAQ92fqbbvuehsLZ3fDEHE2e/q8RH0HVe8O9Ejh1wVTa9MMjd4+nSIt06U+6cDfdPZGLWynPqrMaC4Jtk8U/TRA2EGDK84vK5mrqbim737iXNB8usQ+lGXDlsHVrVDG/9mP8cFW6AoqVRxwN8FixG+xGZFUcnIxAm5kSY6nE86TS+61v6oajTGLbNakc2tozROjRD6uZqzO4QKBgQDs+rNgKD3+fcamN48hXqC7CwQvg1WxyuV0VkK4MKXMbJDwmbxLSku+xh0P9920CjAupcPcsKFAj2FPkEdbkA9Wmm5wTlh9XqauDuKVAkqd77DwkvB4WoY2bQGov16qkQ420uvDc+DysbE7cUvVcl8qcn7YcSO0NsM9+PYlb1GxuQKBgQC7Vij1cDIIrltkx3pggNyiy+nEpx4uhxuDfO8xVSvI5FwSDExPyza50isDKWxJow5iUw9Rz2ybJ2emBeLKwm4G0Cj7Zl/001nsAevu/ob4Hip9qSl+0+Z+Idmvd6BFEYJHoRNK7FjHdwXdBNydwU8BtZZzU3/yHzZ3OIkZFoAvwwKBgQCY6H67advujNuTvr+1CWjup4IQ4k52BPQfJ9WvIXyptdej156+efb75rsz5XyBQh2qy7zgdnvlu15Px0mz0/WBrO3bu0Gvy1YDc4lSGoNo+xMRd85/6fE1xwpOBwUfS69/QoNrvyaDkpJIR6dl14F+UxhzsjUWgEtkfnLc3sI4yQKBgBUMChABKe3lwOjirHIZKDC2Hi505CQwE7xDFhCB1Ch+14VDknNIjn50CVcSmVLwmdYcJNV5K2eHFtMFSESlcX0cd+4+wzsbX7fvQ1WXjQxlPzrc/Yd9QSEcpntbQktgOzXW9/br9NF8ItGBEVQ7+qdjgmK0l+RY82KTnHuQFpjHAoGBAN450ijHPUn67IP9r/KyGrtr21y7Oca1EDpKXbM6p3/H7/637ygf08iJpcYdsh7cslkIVOBxmvPYJZTHlCCD2HoSbEMdwou18rrTst0P72772fayLq6JDvjSjR+fK0blU3RXBlvPAnpi5K/YducknZWitAynC6XZEIytKLDZFTsF"
	bytePrivateKey := FormatPrivateKey(privateKey)
	pri, err := ParsePKCS1PrivateKey(bytePrivateKey)
	if err != nil {
		t.Errorf("ParsePKCS1PrivateKey err:%v\n", err)
		return
	}
	// 公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAma31S+2cqIDZ9jb3kasYR3kVkzCs0r2hw53WX+AqNOWdIRyT762diusA6Xn/EExJ1ph1C9rWBUQy+GaM5blqz1QZbsXfVN3ITJFtWxhPN+jkKqZjdQ1mrGRxwTlcCGFbRy2FhdxsL4x2tsZDVh4P05/+3ZYuVQv2ep8O69DfEVMFbSmXUif73c2eSOthXV8OLN+gar2yhuIAkZoui35cdklfmtV3gNzNO+ykZrIVTzJr4/kIQVlxBRexdFZWgXrE35Kji8ql9p/lF/1v0tXsH0KHqg9L4HIGZ69yEXojdnplZ0h9Is4TRpFlBCDUZFS4AXJqLKQKd+lg1DkEuh6gowIDAQAB"
	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = ParsePKCS1PublicKey(FormatPublicKey(aliPublicKey))
		if err != nil {
			t.Errorf("ParsePKCS1PublicKey err:%v\n", err)
			return
		}
	}
	aliAppClient := &AliAppClient{
		PublicKey:  pub,
		PrivateKey: pri,
		AppID:      "2018030602325750",
	}
	refund := common.AliRefundRequest{
		OutTradeNo:   "AliPay20190801162148WaTu6mgRFT",
		RefundAmount: "0.1",
		RefundReason: "退款测试",
		OutRequestNo: "test",
		OperatorId:   "hcf",
	}
	res, err := aliAppClient.Refund(&refund)
	if err != nil {
		t.Errorf("TestRefund err:%v\n", err)
	}
	t.Errorf("response:%v\n", res)
}

// TestAliTradePay 支付宝统一收单测试
func TestAliTradePay(t *testing.T) {
	// 私钥
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtauMp34N+g+hyrD6845SBaatoU7iSzH/pcyRwaW07z/kREYtpCkpPOTo4O570TR5pexQyhMbg+oi87EnO2dUJkOzkj6rAMcioNZMeXY27AjmbxHy+LwvPSl3tF6nZDU1s+DGEJzM2imdNoa2h117STUs9kwE7c+SE6E3JuDrmqI+4+DkPOJBmgwnN7TVU5AC+Cckv2Uv1shzG1P1P/1/2AzLKqh7ghjgy6uQgsHvU5zaGawS3KtFmVXGiBCekbMrhvo98LRv40D9mYOBsW8t9kLBQMIleNiU4aP+6JzGDRbrnGM0vgYfQKe9C7IK9X/JQEwlpIOI6d6qraCONW1brAgMBAAECggEAEs2eJ/ImTdd7osNuYgjDF20fusYpIzGtRODJOK8VuwCH3wPp+8+z0vc/is1cJN0fyQwhWoDvF4HSxblRH26bHNhr9zRkrUY4nZSBiS9XqMlK+crKQ8zSGP3VRVnlfrVkicY3iD6/3NAQ92fqbbvuehsLZ3fDEHE2e/q8RH0HVe8O9Ejh1wVTa9MMjd4+nSIt06U+6cDfdPZGLWynPqrMaC4Jtk8U/TRA2EGDK84vK5mrqbim737iXNB8usQ+lGXDlsHVrVDG/9mP8cFW6AoqVRxwN8FixG+xGZFUcnIxAm5kSY6nE86TS+61v6oajTGLbNakc2tozROjRD6uZqzO4QKBgQDs+rNgKD3+fcamN48hXqC7CwQvg1WxyuV0VkK4MKXMbJDwmbxLSku+xh0P9920CjAupcPcsKFAj2FPkEdbkA9Wmm5wTlh9XqauDuKVAkqd77DwkvB4WoY2bQGov16qkQ420uvDc+DysbE7cUvVcl8qcn7YcSO0NsM9+PYlb1GxuQKBgQC7Vij1cDIIrltkx3pggNyiy+nEpx4uhxuDfO8xVSvI5FwSDExPyza50isDKWxJow5iUw9Rz2ybJ2emBeLKwm4G0Cj7Zl/001nsAevu/ob4Hip9qSl+0+Z+Idmvd6BFEYJHoRNK7FjHdwXdBNydwU8BtZZzU3/yHzZ3OIkZFoAvwwKBgQCY6H67advujNuTvr+1CWjup4IQ4k52BPQfJ9WvIXyptdej156+efb75rsz5XyBQh2qy7zgdnvlu15Px0mz0/WBrO3bu0Gvy1YDc4lSGoNo+xMRd85/6fE1xwpOBwUfS69/QoNrvyaDkpJIR6dl14F+UxhzsjUWgEtkfnLc3sI4yQKBgBUMChABKe3lwOjirHIZKDC2Hi505CQwE7xDFhCB1Ch+14VDknNIjn50CVcSmVLwmdYcJNV5K2eHFtMFSESlcX0cd+4+wzsbX7fvQ1WXjQxlPzrc/Yd9QSEcpntbQktgOzXW9/br9NF8ItGBEVQ7+qdjgmK0l+RY82KTnHuQFpjHAoGBAN450ijHPUn67IP9r/KyGrtr21y7Oca1EDpKXbM6p3/H7/637ygf08iJpcYdsh7cslkIVOBxmvPYJZTHlCCD2HoSbEMdwou18rrTst0P72772fayLq6JDvjSjR+fK0blU3RXBlvPAnpi5K/YducknZWitAynC6XZEIytKLDZFTsF"
	bytePrivateKey := FormatPrivateKey(privateKey)
	pri, err := ParsePKCS1PrivateKey(bytePrivateKey)
	if err != nil {
		t.Errorf("ParsePKCS1PrivateKey err:%v\n", err)
		return
	}
	// 公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAma31S+2cqIDZ9jb3kasYR3kVkzCs0r2hw53WX+AqNOWdIRyT762diusA6Xn/EExJ1ph1C9rWBUQy+GaM5blqz1QZbsXfVN3ITJFtWxhPN+jkKqZjdQ1mrGRxwTlcCGFbRy2FhdxsL4x2tsZDVh4P05/+3ZYuVQv2ep8O69DfEVMFbSmXUif73c2eSOthXV8OLN+gar2yhuIAkZoui35cdklfmtV3gNzNO+ykZrIVTzJr4/kIQVlxBRexdFZWgXrE35Kji8ql9p/lF/1v0tXsH0KHqg9L4HIGZ69yEXojdnplZ0h9Is4TRpFlBCDUZFS4AXJqLKQKd+lg1DkEuh6gowIDAQAB"
	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = ParsePKCS1PublicKey(FormatPublicKey(aliPublicKey))
		if err != nil {
			t.Errorf("ParsePKCS1PublicKey err:%v\n", err)
			return
		}
	}
	aliAppClient := &AliAppClient{
		PublicKey:  pub,
		PrivateKey: pri,
		AppID:      "2018030602325750",
	}
	req := common.AliTradePayRequest{
		OutTradeNo:    "AliPay201955edd62166WaTu6mgRFT",
		Scene:         "bar_code",
		AuthCode:      "288158987026118204",
		Subject:       "条码付款测试",
		TotalAmount:   "0.01",
		TransCurrency: "CNY",
	}
	res, err := aliAppClient.AliTradePay(&req)
	if err != nil {
		t.Errorf("TestAliTradePay err:%v\n", err)
	}
	t.Errorf("response:%v\n", res)
}

// TestAliQueryOrder 支付宝订单查询接口
func TestAliQueryOrder(t *testing.T) {
	// 私钥
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtauMp34N+g+hyrD6845SBaatoU7iSzH/pcyRwaW07z/kREYtpCkpPOTo4O570TR5pexQyhMbg+oi87EnO2dUJkOzkj6rAMcioNZMeXY27AjmbxHy+LwvPSl3tF6nZDU1s+DGEJzM2imdNoa2h117STUs9kwE7c+SE6E3JuDrmqI+4+DkPOJBmgwnN7TVU5AC+Cckv2Uv1shzG1P1P/1/2AzLKqh7ghjgy6uQgsHvU5zaGawS3KtFmVXGiBCekbMrhvo98LRv40D9mYOBsW8t9kLBQMIleNiU4aP+6JzGDRbrnGM0vgYfQKe9C7IK9X/JQEwlpIOI6d6qraCONW1brAgMBAAECggEAEs2eJ/ImTdd7osNuYgjDF20fusYpIzGtRODJOK8VuwCH3wPp+8+z0vc/is1cJN0fyQwhWoDvF4HSxblRH26bHNhr9zRkrUY4nZSBiS9XqMlK+crKQ8zSGP3VRVnlfrVkicY3iD6/3NAQ92fqbbvuehsLZ3fDEHE2e/q8RH0HVe8O9Ejh1wVTa9MMjd4+nSIt06U+6cDfdPZGLWynPqrMaC4Jtk8U/TRA2EGDK84vK5mrqbim737iXNB8usQ+lGXDlsHVrVDG/9mP8cFW6AoqVRxwN8FixG+xGZFUcnIxAm5kSY6nE86TS+61v6oajTGLbNakc2tozROjRD6uZqzO4QKBgQDs+rNgKD3+fcamN48hXqC7CwQvg1WxyuV0VkK4MKXMbJDwmbxLSku+xh0P9920CjAupcPcsKFAj2FPkEdbkA9Wmm5wTlh9XqauDuKVAkqd77DwkvB4WoY2bQGov16qkQ420uvDc+DysbE7cUvVcl8qcn7YcSO0NsM9+PYlb1GxuQKBgQC7Vij1cDIIrltkx3pggNyiy+nEpx4uhxuDfO8xVSvI5FwSDExPyza50isDKWxJow5iUw9Rz2ybJ2emBeLKwm4G0Cj7Zl/001nsAevu/ob4Hip9qSl+0+Z+Idmvd6BFEYJHoRNK7FjHdwXdBNydwU8BtZZzU3/yHzZ3OIkZFoAvwwKBgQCY6H67advujNuTvr+1CWjup4IQ4k52BPQfJ9WvIXyptdej156+efb75rsz5XyBQh2qy7zgdnvlu15Px0mz0/WBrO3bu0Gvy1YDc4lSGoNo+xMRd85/6fE1xwpOBwUfS69/QoNrvyaDkpJIR6dl14F+UxhzsjUWgEtkfnLc3sI4yQKBgBUMChABKe3lwOjirHIZKDC2Hi505CQwE7xDFhCB1Ch+14VDknNIjn50CVcSmVLwmdYcJNV5K2eHFtMFSESlcX0cd+4+wzsbX7fvQ1WXjQxlPzrc/Yd9QSEcpntbQktgOzXW9/br9NF8ItGBEVQ7+qdjgmK0l+RY82KTnHuQFpjHAoGBAN450ijHPUn67IP9r/KyGrtr21y7Oca1EDpKXbM6p3/H7/637ygf08iJpcYdsh7cslkIVOBxmvPYJZTHlCCD2HoSbEMdwou18rrTst0P72772fayLq6JDvjSjR+fK0blU3RXBlvPAnpi5K/YducknZWitAynC6XZEIytKLDZFTsF"
	bytePrivateKey := FormatPrivateKey(privateKey)
	pri, err := ParsePKCS1PrivateKey(bytePrivateKey)
	if err != nil {
		t.Errorf("ParsePKCS1PrivateKey err:%v\n", err)
		return
	}
	// 公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAma31S+2cqIDZ9jb3kasYR3kVkzCs0r2hw53WX+AqNOWdIRyT762diusA6Xn/EExJ1ph1C9rWBUQy+GaM5blqz1QZbsXfVN3ITJFtWxhPN+jkKqZjdQ1mrGRxwTlcCGFbRy2FhdxsL4x2tsZDVh4P05/+3ZYuVQv2ep8O69DfEVMFbSmXUif73c2eSOthXV8OLN+gar2yhuIAkZoui35cdklfmtV3gNzNO+ykZrIVTzJr4/kIQVlxBRexdFZWgXrE35Kji8ql9p/lF/1v0tXsH0KHqg9L4HIGZ69yEXojdnplZ0h9Is4TRpFlBCDUZFS4AXJqLKQKd+lg1DkEuh6gowIDAQAB"
	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = ParsePKCS1PublicKey(FormatPublicKey(aliPublicKey))
		if err != nil {
			t.Errorf("ParsePKCS1PublicKey err:%v\n", err)
			return
		}
	}
	aliAppClient := &AliAppClient{
		PublicKey:  pub,
		PrivateKey: pri,
		AppID:      "2018030602325750",
	}
	res, err := aliAppClient.QueryOrder("AliPay201955e1162166WaTu6mgRFT")
	if err != nil {
		t.Errorf("QueryOrder err:%v\n", err)
	}
	t.Errorf("response:%v\n", res)
}

// TestAliTradeCancel 支付宝撤单测试
func TestAliTradeCancel(t *testing.T) {
	// 私钥
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCtauMp34N+g+hyrD6845SBaatoU7iSzH/pcyRwaW07z/kREYtpCkpPOTo4O570TR5pexQyhMbg+oi87EnO2dUJkOzkj6rAMcioNZMeXY27AjmbxHy+LwvPSl3tF6nZDU1s+DGEJzM2imdNoa2h117STUs9kwE7c+SE6E3JuDrmqI+4+DkPOJBmgwnN7TVU5AC+Cckv2Uv1shzG1P1P/1/2AzLKqh7ghjgy6uQgsHvU5zaGawS3KtFmVXGiBCekbMrhvo98LRv40D9mYOBsW8t9kLBQMIleNiU4aP+6JzGDRbrnGM0vgYfQKe9C7IK9X/JQEwlpIOI6d6qraCONW1brAgMBAAECggEAEs2eJ/ImTdd7osNuYgjDF20fusYpIzGtRODJOK8VuwCH3wPp+8+z0vc/is1cJN0fyQwhWoDvF4HSxblRH26bHNhr9zRkrUY4nZSBiS9XqMlK+crKQ8zSGP3VRVnlfrVkicY3iD6/3NAQ92fqbbvuehsLZ3fDEHE2e/q8RH0HVe8O9Ejh1wVTa9MMjd4+nSIt06U+6cDfdPZGLWynPqrMaC4Jtk8U/TRA2EGDK84vK5mrqbim737iXNB8usQ+lGXDlsHVrVDG/9mP8cFW6AoqVRxwN8FixG+xGZFUcnIxAm5kSY6nE86TS+61v6oajTGLbNakc2tozROjRD6uZqzO4QKBgQDs+rNgKD3+fcamN48hXqC7CwQvg1WxyuV0VkK4MKXMbJDwmbxLSku+xh0P9920CjAupcPcsKFAj2FPkEdbkA9Wmm5wTlh9XqauDuKVAkqd77DwkvB4WoY2bQGov16qkQ420uvDc+DysbE7cUvVcl8qcn7YcSO0NsM9+PYlb1GxuQKBgQC7Vij1cDIIrltkx3pggNyiy+nEpx4uhxuDfO8xVSvI5FwSDExPyza50isDKWxJow5iUw9Rz2ybJ2emBeLKwm4G0Cj7Zl/001nsAevu/ob4Hip9qSl+0+Z+Idmvd6BFEYJHoRNK7FjHdwXdBNydwU8BtZZzU3/yHzZ3OIkZFoAvwwKBgQCY6H67advujNuTvr+1CWjup4IQ4k52BPQfJ9WvIXyptdej156+efb75rsz5XyBQh2qy7zgdnvlu15Px0mz0/WBrO3bu0Gvy1YDc4lSGoNo+xMRd85/6fE1xwpOBwUfS69/QoNrvyaDkpJIR6dl14F+UxhzsjUWgEtkfnLc3sI4yQKBgBUMChABKe3lwOjirHIZKDC2Hi505CQwE7xDFhCB1Ch+14VDknNIjn50CVcSmVLwmdYcJNV5K2eHFtMFSESlcX0cd+4+wzsbX7fvQ1WXjQxlPzrc/Yd9QSEcpntbQktgOzXW9/br9NF8ItGBEVQ7+qdjgmK0l+RY82KTnHuQFpjHAoGBAN450ijHPUn67IP9r/KyGrtr21y7Oca1EDpKXbM6p3/H7/637ygf08iJpcYdsh7cslkIVOBxmvPYJZTHlCCD2HoSbEMdwou18rrTst0P72772fayLq6JDvjSjR+fK0blU3RXBlvPAnpi5K/YducknZWitAynC6XZEIytKLDZFTsF"
	bytePrivateKey := FormatPrivateKey(privateKey)
	pri, err := ParsePKCS1PrivateKey(bytePrivateKey)
	if err != nil {
		t.Errorf("ParsePKCS1PrivateKey err:%v\n", err)
		return
	}
	// 公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAma31S+2cqIDZ9jb3kasYR3kVkzCs0r2hw53WX+AqNOWdIRyT762diusA6Xn/EExJ1ph1C9rWBUQy+GaM5blqz1QZbsXfVN3ITJFtWxhPN+jkKqZjdQ1mrGRxwTlcCGFbRy2FhdxsL4x2tsZDVh4P05/+3ZYuVQv2ep8O69DfEVMFbSmXUif73c2eSOthXV8OLN+gar2yhuIAkZoui35cdklfmtV3gNzNO+ykZrIVTzJr4/kIQVlxBRexdFZWgXrE35Kji8ql9p/lF/1v0tXsH0KHqg9L4HIGZ69yEXojdnplZ0h9Is4TRpFlBCDUZFS4AXJqLKQKd+lg1DkEuh6gowIDAQAB"
	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = ParsePKCS1PublicKey(FormatPublicKey(aliPublicKey))
		if err != nil {
			t.Errorf("ParsePKCS1PublicKey err:%v\n", err)
			return
		}
	}
	aliAppClient := &AliAppClient{
		PublicKey:  pub,
		PrivateKey: pri,
		AppID:      "2018030602325750",
	}
	req := common.AliTradeCancelRequest{
		OutTradeNo: "AliPay201955e1162166WaTu6mgRFT",
	}
	res, err := aliAppClient.AliTradeCancel(&req)
	if err != nil {
		t.Errorf("AliTradeCancel err:%v\n", err)
	}
	t.Errorf("response:%v\n", res)
}
