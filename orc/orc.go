package ocr

import (
	"encoding/json"
	"errors"
	"github.com/jxwt/tools"
)

const apiBusinessLicense = "https//dm-58.data.aliyun.com/rest/160601/ocr/ocr_business_license.json"

type BusinessLicenseRequest struct {
	Image string
}

type BusinessLicenseResponse struct {
	Angle         float64 //输入图片的角度（顺时针旋转），［0， 90， 180，270］
	RegNm         string  `json:"reg_nm"` //注册号，没有识别出来时返回FailInRecognition
	Name          string  //公司名称，没有识别出来时返回FailInRecognition
	Person        string  //公司法人，没有识别出来时返回FailInRecognition
	EstablishDate string  `json:"establish_date"` //公司注册日期(例：证件上为2014年04月16日，算法返回20140416)
	ValidPeriod   string  `json:"valid_period"`   //公司营业期限终止日期(例：证件上为2014年04月16日至2034年04月15日，算法返回20340415)
	//当前算法将日期格式统一为输出为年月日(如20391130)并将长期表示为29991231，若证件上没有营业期限，则默认其为长期返回29991231。
	Address   string //公司地址，没有识别出来时返回FailInRecognition
	Captial   string //注册资本，没有识别出来时返回FailInRecognition
	Business  string //经营范围，没有识别出来时返回FailInRecognition
	Elbem     string //国徽位置［topleftheightwidth］，没有识别出来时返回FailInDetection
	Title     string //标题位置［topleftheightwidth］，没有识别出来时返回FailInDetection
	Stamp     string //印章位置［topleftheightwidth］，没有识别出来时返回FailInDetection
	Qrcode    string //二维码位置［topleftheightwidth］，没有识别出来时返回FailInDetection
	success   bool   //识别成功与否 true/false
	RequestId string
}

func BusinessLicense(url string) (*BusinessLicenseResponse, error) {
	req := &BusinessLicenseRequest{
		Image: url,
	}
	data, _ := json.Marshal(req)
	body, err := tools.HttpBeegoJsonPost(apiBusinessLicense, string(data), nil)
	if err != nil {
		return nil, err
	}
	resp := new(BusinessLicenseResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	if resp.success {
		return resp, nil
	}
	return nil, errors.New("识别错误, 请确认营业执照图片")
}
