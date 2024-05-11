package db_test

import (
	"intellipath/internal/record"
	"testing"
)


func TestGetDB(t *testing.T) {
	test_db, _ := record.GetDatabase("test_paths.db")
	if test_db == nil {
		t.FailNow()
	}
	
}

func TestDBInsertion(t *testing.T) {
	test_db, _ := record.GetDatabase("test_paths.db")

	paths, err := test_db.GetAllPaths()
	if err != nil {
		t.FailNow()
	}

	numOfRecs := len(paths)

	rec, _ := record.NewRecord("/home/nirt", 0)
	_, err = test_db.InsertRecord(rec)
	
	if err != nil {
		t.FailNow()
	}

	paths, _ = test_db.GetAllPaths()

	if numOfRecs + 1 !=  len(paths){
		t.FailNow()
	}
}

func TestGetRecordsByName(t *testing.T) {
	test_db, _ := record.GetDatabase("test_paths.db")
	if test_db == nil {
		t.FailNow()
	}

	paths := []string{"/home/nirt"}
	rec, err := test_db.GetRecordsByName(paths)

	if err != nil {
		t.FailNow()
	}

	if len(rec) != 1 {
		t.FailNow()
	}

}
