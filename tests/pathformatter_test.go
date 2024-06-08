package internal_test

import (
	"intellipath/internal/utils"
	"testing"
)

func TestToBase(t *testing.T) {
	formatter := utils.NewPathFormatter()
	base := formatter.ToBase("/home/user/Desktop")
	if base != "Desktop" {
		t.Errorf("Formatter - ToBase function failed")
	}
}

func TestIsExists(t *testing.T) {
	formatter := utils.NewPathFormatter()
	isExists := formatter.IsExists("/home/nirt/Desktop")
	if isExists != true {
		t.Errorf("Formatter - IsExist function failed")
	}
}
