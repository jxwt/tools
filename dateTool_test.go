package tools

import (
	"testing"
	"time"
)

func TestGetFirstDateOfMonth(t *testing.T) {
	t.Logf("%v", GetFirstDateOfMonth(time.Now()))
	t.Logf("%v", GetLastDateOfMonth(time.Now()))
}
