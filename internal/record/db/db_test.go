package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// var DB *SQLDatabase

func setup() (*SQLDatabase,error) {
	// Create a new database file for testing
	dbPath := "../../testsdata/test.db"

	// Remove existing database file if it exists
	abs,_ := filepath.Abs(dbPath)
	if _, err := os.Stat(abs); err == nil {
		os.Remove(abs)
	}

	os.Create(abs)
	// Initialize database and perform necessary setup
	DB, err := GetDBInstance(abs)
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance: %v", err)
	}

	err = DB.Initizlize()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DB: %v", err)
	}

	return DB, nil
}

func teardown() {
	// Remove database file after test completes
	dbPath := "../../testsdata/test.db"
	abs,_ := filepath.Abs(dbPath)

	os.Remove(abs)
}

func TestInitialize(t *testing.T) {
	DB, err := setup()
	assert.NoError(t, err, "Setup should not return an error")
	defer teardown()

	// Check for table existence
	rows, err := DB.db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='paths'")
	assert.NoError(t, err)
	defer rows.Close()
	assert.True(t, rows.Next(), "Table 'paths' should exist")

	// Check for trigger existence
	rows, err = DB.db.Query("SELECT name FROM sqlite_master WHERE type='trigger' AND name='update_last_touched_after_insert'")
	assert.NoError(t, err)
	defer rows.Close()
	assert.True(t, rows.Next(), "Trigger 'update_last_touched_after_insert' should exist")

	// Check for another trigger existence
	rows, err = DB.db.Query("SELECT name FROM sqlite_master WHERE type='trigger' AND name='update_last_touched_after_update'")
	assert.NoError(t, err)
	defer rows.Close()
	assert.True(t, rows.Next(), "Trigger 'update_last_touched_after_update' should exist")
}