package db_test

import (
	"intellipath/internal/record"
	"testing"
)

var testDB *record.Database

func setup() {
	db := record.GetDbInstance("file::memory:?cache=shared")
	db.Initizlize()

	testDB = db
}

func teardown(db *record.Database) {
	db.Close()
}

func TestDBInsertion(t *testing.T) {
	setup()

	paths, err := testDB.GetAllPaths()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	numOfRecs := len(paths)

	rec, _ := record.NewRecord("/home/nirt", 0)
	_, err = testDB.InsertRecord(rec)
	
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	paths, _ = testDB.GetAllPaths()
	if numOfRecs + 1 !=  len(paths){
		t.FailNow()
	}
}

func TestGetRecordsByName(t *testing.T) {
	setup()

	new_rec, _ := record.NewRecord("/home/nirt/Desktop", 0)
	_, err := testDB.InsertRecord(new_rec)
	if err != nil {
		t.Log(err)
	}
	paths := []string{"/home/nirt/Desktop"}
	rec, err := testDB.GetRecordsByName(paths)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if len(rec) != 1 {
		t.FailNow()
	}
}

func TestMain(m *testing.M) {
	m.Run()
	defer teardown(testDB)
}