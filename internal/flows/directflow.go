package pathfinder

import (
	"intellipath/internal/record"
	"intellipath/internal/record/db"
)

type Direct struct {
	absolutePath string
	db db.Database
}

func NewDirectFlow(absolutePath string, db db.Database) *Direct {
	return &Direct{
		absolutePath: absolutePath,
		db: db,
	}
}

func (direct *Direct) FindMatch() string { // This should later on return a record
	var outPath string

	rec, err := direct.db.PathSearch(direct.absolutePath) // This should return a record if it exists
	if err != nil {
		return ""
	}

	switch rec.GetPath() {
	case "": // In case no record was found
		record, err := record.NewRecord(direct.absolutePath, 0)
		if err != nil {
			return ""
		}

		direct.SaveRecord(record)
		// if _, err = direct.db.InsertRecord(record); err != nil {
		// 	return ""
		// }
		outPath = direct.absolutePath

	case direct.absolutePath: // In case a matching record was found
		direct.UpdateRecord(rec)
		// if err := direct.db.UpdateScore(rec); err != nil {
		// 	return ""
		// }
		outPath = direct.absolutePath

	}

	return outPath
}

func (direct *Direct) SaveRecord(record *record.PathRecord) {
	_, err := direct.db.InsertRecord(record)
	if err != nil {
		panic(err)
	}
}

func (direct *Direct) UpdateRecord(record *record.PathRecord) {
	if err := direct.db.UpdateScore(record); err != nil {
		panic(err)
	}
}