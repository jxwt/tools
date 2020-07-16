package tools

import "testing"

func TestZip(t *testing.T){
	srcDir:="/Users/hcf/Pictures/test_pic"
	zipFileName:="test.zip"
	Zip(srcDir,zipFileName)
}
