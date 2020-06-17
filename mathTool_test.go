package tools

import (
	"testing"
)

func TestIntToStringLength(t *testing.T) {
	a := IntToStringLength(10, 3)
	t.Logf("%v", a)
	b := IntToStringLength(109, 3)
	t.Logf("%v", b)
	c := IntToStringLength(1091, 3)
	t.Logf("%v", c)
}
