package pathfinder

import (
	"errors"
	"intellipath/internal/record"
	"intellipath/internal/record/db"
	"intellipath/internal/matcher"

)

type FuzzyFlow struct {
	basePath string
	db db.Database
}

func NewFuzzyFlow(basePath string, db db.Database) *FuzzyFlow {
	return &FuzzyFlow{
		basePath: basePath,
		db: db,
	}
}

func (fuzzy *FuzzyFlow) FindMatch() string {
	// Check in DB + fuzzy + levinshtein
	pathFormatter := record.NewPathFormatter()

	records := fuzzy.FetchRecords()

	foundRecords, err := matcher.FuzzyFind(fuzzy.basePath, records) //fuzzy + levinshtein getting a PathDistRecord struct
	if err != nil {
		return ""
	}

	var fuzzyPaths []string
	for _, pathRes := range foundRecords {
		fuzzyPaths = append(fuzzyPaths, pathRes.Path)
	}

	records, err = fuzzy.db.GetRecordsByName(fuzzyPaths)
	if err != nil {
		return ""
	}

	// filter by score function
	rec, err := fuzzy.filterByScore(records)
	if err != nil {
		return ""
	}

	// Check if found path exists
	if !pathFormatter.IsExists(rec.GetPath()) {
		fuzzy.db.DeletePath(rec)
		panic("Path does not exists")
	} else {
		_ = fuzzy.db.UpdateScore(rec)
		return rec.GetPath()
	}
}

func (fuzzy *FuzzyFlow) filterByScore(records []*record.PathRecord) (*record.PathRecord, error) {

	if len(records) < 1 {
		return nil, errors.New("could not find any paths")
	} else if len(records) == 1 {
		return records[0], nil
	} else {
		if records[0].GetScore() > records[1].GetScore() {
			return records[0], nil
		} else {
			return records[1], nil
		}
	}
}

func  (fuzzy *FuzzyFlow) FetchRecords() []*record.PathRecord {
	records, err := fuzzy.db.GetAllRecords();
	if err != nil {
		panic(err)
	}

	return records
}

func  (fuzzy *FuzzyFlow) SaveRecord() {}