package tools

import "testing"

// TestPollQuery .
func TestPollQuery(t *testing.T) {
	client := &Kuaidi100Sender{
		Customer: "*",
		Key:      "*",
	}
	res, err := client.PollQuery("zhongtong", "*")
	//err := client.PollQuery("huitongkuaidi", "*")
	t.Logf("%v", res)
	t.Logf("%v", err)
}
