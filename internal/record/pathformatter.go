package record

import (
	"os"
	"path/filepath"
)

// type Formatter interface {
// 	ToBase(string) string
// 	ToAbs(string) string
// 	IsExists(string) bool
// }

type PathFormatter struct{}


func NewPathFormatter() *PathFormatter {
	return &PathFormatter{}
}

func (f *PathFormatter) ToBase(path string) string {
	return filepath.Base(path)
}

func (f *PathFormatter) ToAbs(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic("Could not convert to absolute path")
	}

	return absPath
}

func (f *PathFormatter) IsExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
