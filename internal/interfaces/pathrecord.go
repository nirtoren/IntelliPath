// Redundant currently


package interfaces

import (
	"intellipath/internal/db"
)

type PathRecord struct{
	Record
	Path string
	Score int8
}

type Record interface{
	ScoreUpdate(*db.Database, db.PathRecord) error
	GetScore(db.PathRecord) int
}

func ScoreUpdate(db *db.Database, record db.PathRecord) error {
	return db.UpdateScore(record.Path, record.Score)
}

func GetScore(record db.PathRecord) int {
	return int(record.Score)
}

func GetPath(record db.PathRecord) string {
	return record.Path
}

