package tools

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Kuaidi100Sender struct {
	Customer string
	Key      string
}

const (
	// PollQueryURL 实时查询接口
	PollQueryURL = "https://poll.kuaidi100.com/poll/query.do"
)

// PollQueryRequest 实时查询接口请求
type PollQueryRequest struct {
	Com string `json:"com"` // 快递公司
	Num string `json:"num"` // 快递单号
}

// PollQueryResponse 返回
type PollQueryResponse struct {
	Message   string `json:"message"`
	Nu        string `json:"nu"`
	Ischeck   string `json:"ischeck"`
	Condition string `json:"condition"`
	Com       string `json:"com"`
	Status    string `json:"status"`
	State     string `json:"state"`
	Data      []struct {
		Time    string `json:"time"`
		Ftime   string `json:"ftime"`
		Context string `json:"context"`
	} `json:"data"`
}

// PollQuery 实时查询接口
func (c *Kuaidi100Sender) PollQuery(com string, num string) (*PollQueryResponse, error) {
	req := PollQueryRequest{
		Com: com,
		Num: num,
	}
	param, err := json.Marshal(req)
	if err != nil {
		logs.Warning(err)
		return nil, err
	}
	sign := SignMD5Up(string(param) + c.Key + c.Customer)
	customer := c.Customer
	values := &url.Values{}
	values.Set("customer", customer)
	values.Set("sign", sign)
	values.Set("param", string(param))
	resp, err := http.PostForm(PollQueryURL, *values)
	if err != nil {
		logs.Warning(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Warning(err)
		logs.Warning(string(body))
		return nil, err
	}
	var res PollQueryResponse
	if err := json.Unmarshal(body, &res); err != nil {
		logs.Warning(err)
		return nil, err
	}
	return &res, err
}
