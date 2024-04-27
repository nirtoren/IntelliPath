package interfaces

import (
	"os"
	"path/filepath"
)

type Formatter interface {
	ToBase(string) string
	ToAbs(string) string
	IsExists(string) bool
}

func ToBase(path string) string {
	return filepath.Base(path)
}

func ToAbs(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic("Could not convert to absolute path")
	}

	return absPath
}

func IsExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
