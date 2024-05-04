package pathformatter_test

import (
	"testing"
	"intellipath/internal/interfaces"
)

func TestToBase(t *testing.T) {
	formatter := interfaces.NewPathFormatter()
	base := formatter.ToBase("/home/user/Desktop")
	if base != "Desktop" {
		t.Errorf("Formatter - ToBase function failed")
	}
}

func TestIsExists(t *testing.T) {
	formatter := interfaces.NewPathFormatter()
	isExists := formatter.IsExists("/home/nirt/Desktop")
	if isExists != true {
		t.Errorf("Formatter - IsExist function failed")
	}
}
