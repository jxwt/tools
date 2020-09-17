package tools

import (
	"testing"
)

func TestUploadFileObs(t *testing.T) {
	InitObs("GTKY0N2IZGQXIN5ZSPKE", "aq12j84UvglDBkfuskqWMYoq4QKG7aKFy6EpMzBi", "obs.cn-east-3.myhuaweicloud.com", "jx-test")
	//
	//f, err := os.Open("20200915101047_03.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//ret, err := UploadFileObs(time.Now().Format("20060102150405")+ GetRandomString(12) + ".jpg", f)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ret)
}
