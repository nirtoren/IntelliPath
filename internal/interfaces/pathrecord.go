// Redundant currently

package interfaces

import (
	"errors"
)

type Record interface{
	GetScore() int
	GetPath() string
}

type PathRecord struct{
	path string
	score int
}


func NewRecord(path string, score int) (*PathRecord, error) {
	if path == "" {
		return nil, errors.New("path can not be NULL")
	}
	return &PathRecord{path: path,
		score: score}, nil
}

func (r *PathRecord) GetScore() int {
	return int(r.score)
}

func (r *PathRecord) GetPath() string {
	return r.path
}

