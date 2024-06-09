package record

import (
	"testing"
)

func TestToBase(t *testing.T) {
	formatter := NewPathFormatter()
	base := formatter.ToBase("/home/user/Desktop")
	if base != "Desktop" {
		t.Errorf("Formatter - ToBase function failed")
	}
}

func TestIsExists(t *testing.T) {
	formatter := NewPathFormatter()
	isExists := formatter.IsExists("/home/nirt/Desktop")
	if isExists != true {
		t.Errorf("Formatter - IsExist function failed")
	}
}
