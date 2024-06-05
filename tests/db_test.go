package internal_test

import (
	"intellipath/internal/record"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock PathFormatter
type PathFormatter interface {
	ToAbs(path string) string
	IsExists(path string) bool
}

type MockPathFormatter struct {
	mock.Mock
}

func (m *MockPathFormatter) ToAbs(path string) string {
	args := m.Called(path)
	return args.String(0)
}

func (m *MockPathFormatter) IsExists(path string) bool {
	args := m.Called(path)
	return args.Bool(0)
}

var testDBFile = "test.db"

func setup() {
	// Ensure the test database file does not exist before starting tests
	if _, err := os.Stat(testDBFile); err == nil {
		os.Remove(testDBFile)
	}
	
	f,err := os.Create(testDBFile)
	if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
	
}

func teardown() {
	// Remove the test database file after tests
	os.Remove(testDBFile)
}


func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGetDbInstance(t *testing.T) {
	mockFormatter := new(MockPathFormatter)
	dbFile := "test.db"

	t.Run("Database exists", func(t *testing.T) {
		mockFormatter.On("ToAbs", dbFile).Return(dbFile)
		mockFormatter.On("IsExists", dbFile).Return(true)

		db := record.GetDbInstance(dbFile)
		assert.NotNil(t, db)
		// mockFormatter.AssertExpectations(t)
	})

	t.Run("Database does not exist", func(t *testing.T) {
		os.Remove(testDBFile)
		mockFormatter.On("ToAbs", dbFile).Return(dbFile)
		mockFormatter.On("IsExists", dbFile).Return(false)

		assert.PanicsWithValue(t, "Could not find the database", func() {
			record.GetDbInstance(dbFile)
		})
		// mockFormatter.AssertExpectations(t)
	})
}

func TestDatabase_Initialize(t *testing.T) {
	setup()
	mockFormatter := new(MockPathFormatter)
	mockFormatter.On("ToAbs", "test.db").Return("test.db")
	mockFormatter.On("IsExists", "test.db").Return(true)

	db := record.GetDbInstance("test.db")
	defer db.Close()

	err := db.Initizlize()
	assert.NoError(t, err)
}

func TestDatabase_Close(t *testing.T) {
	mockFormatter := new(MockPathFormatter)
	mockFormatter.On("ToAbs", "test.db").Return("test.db")
	mockFormatter.On("IsExists", "test.db").Return(true)

	db := record.GetDbInstance("test.db")
	err := db.Close()
	assert.NoError(t, err)
}