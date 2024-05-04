package db_test

import (
	"testing"
	"intellipath/internal/db"
)

func TestGetDB(t *testing.T) {
	test_db, _ := db.GetDatabase("test_paths.db")
	if test_db == nil {
		t.FailNow()
	}
	
}

func TestDBInsertion(t *testing.T) {
	test_db, _ := db.GetDatabase("test_paths.db")

	paths, err := test_db.GetAllPaths()
	if err != nil {
		t.FailNow()
	}

	numOfRecs := len(paths)

	rec, _ := db.NewRecord("/home/nirt", 0)
	_, err = test_db.InsertRecord(rec)
	
	if err != nil {
		t.FailNow()
	}

	paths, _ = test_db.GetAllPaths()

	if numOfRecs !=  len(paths) + 1 {
		t.FailNow()
	}
}

func TestGetRecordsByName(t *testing.T) {
	test_db, _ := db.GetDatabase("test_paths.db")
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

func TestUpdateScore(t *testing.T) {
	test_db, _ := db.GetDatabase("test_paths.db")
	if test_db == nil {
		t.FailNow()
	}

	paths := []string{"/home/nirt"}
	rec, _ := test_db.GetRecordsByName(paths)
	
	err := test_db.UpdateScore("/home/nirt", 0)
}