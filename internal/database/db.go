package database

import (
	"database/sql"
	"errors"
	"fmt"
	"intellipath/internal/record"
	"intellipath/internal/utils"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Close() error
	DeletePath(pathRec *record.PathRecord) error
	GetAllRecords() ([]*record.PathRecord, error)
	GetRecordsByName(optionalPaths []string) ([]*record.PathRecord, error)
	Initizlize() error
	InsertRecord(pathRec *record.PathRecord) (int64, error)
	PathSearch(pathToSearch string) (*record.PathRecord, error)
	UpdateScore(pathRec *record.PathRecord) error
}

type SQLDatabase struct {
	db *sql.DB
}

var (
	dbInstance *SQLDatabase
	once       sync.Once
)

func GetDbInstance() *SQLDatabase {
	var dbFile string
	validator := utils.NewValidator()
	ENVGetter := utils.NewENVGetter(validator)
	dbFile = ENVGetter.GetDBPath()

	once.Do(func() {
		var err error
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			panic(err)
		}
		dbInstance = &SQLDatabase{db: db}
	})

	return dbInstance
}


func (d *SQLDatabase) Initizlize() error {
	createSchemaSQL := `
		CREATE TABLE IF NOT EXISTS paths (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL UNIQUE,
			score INTEGER,
			last_touched DATETIME
		)
	`

	_, err := d.db.Exec(createSchemaSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	// trigger for updating 'last_touched' after INSERT
	_, err = d.db.Exec(`
		CREATE TRIGGER IF NOT EXISTS update_last_touched_insert
		AFTER INSERT ON paths
		BEGIN
			UPDATE paths SET last_touched = CURRENT_TIMESTAMP WHERE id = NEW.id;
		END;
	`)
	if err != nil {
		return fmt.Errorf("could not initialize trigger")
	}

	// trigger for updating 'last_touched' after UPDATE
	_, err = d.db.Exec(`
		CREATE TRIGGER IF NOT EXISTS update_last_touched_insert
		AFTER UPDATE ON paths
		BEGIN
			UPDATE paths SET last_touched = CURRENT_TIMESTAMP WHERE id = NEW.id;
		END;
	`)

	if err != nil {
		return fmt.Errorf("could not initialize trigger")
	}

	return nil
}

func (d *SQLDatabase) Close() error {
	if d.db != nil {
		return d.db.Close()
	}

	return nil
}

func ParallelCleanUp(d *SQLDatabase, dtimer int, resultCh chan<- error) {
	cutOffTime := time.Now().Add(-5 * 24 * time.Hour)

	_, err := d.db.Exec("DELETE FROM paths WHERE last_touched < ?", cutOffTime)
	resultCh <- err
}

func (d *SQLDatabase) InsertRecord(pathRec *record.PathRecord) (int64, error) {

	insertPathSQL := `
		INSERT INTO paths (path, score) VALUES (?, ?)
	`
	result, err := d.db.Exec(insertPathSQL, pathRec.GetPath(), pathRec.GetScore())

	if err != nil {
		return 0, fmt.Errorf("error on insertion")
	}

	userID, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error while getting last inserted id: %v", err)
	}

	return userID, nil
}

func (d *SQLDatabase) GetAllRecords() ([]*record.PathRecord, error) {

	var records []*record.PathRecord

	getAllPathsSQL := `
		SELECT path, score FROM paths
	`
	rows, err := d.db.Query(getAllPathsSQL)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var path string
		var score int
		if err := rows.Scan(&path, &score); err != nil {
			return nil, fmt.Errorf("error scanning row %v", err)
		}
		record, _ := record.NewRecord(path, score)
		records = append(records, record)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	return records, nil
}

func (d *SQLDatabase) PathSearch(pathToSearch string) (*record.PathRecord, error) {

	var path string = ""
	var score int = 0

	searchPathsSQL := "SELECT path,score FROM paths WHERE path = ?"
	rows, err := d.db.Query(searchPathsSQL, pathToSearch)
	if err != nil {
		fmt.Println(err)

	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&path, &score); err != nil {
			fmt.Println(err)
		}
	}

	newRec, _ := record.NewRecord(path, score)

	return newRec, err
}

func (d *SQLDatabase) UpdateScore(pathRec *record.PathRecord) error {

	updateScoresSQL := `UPDATE paths SET score = ? WHERE path = ?`
	newScore := pathRec.GetScore() + 1
	_, err := d.db.Exec(updateScoresSQL, newScore, pathRec.GetPath())
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *SQLDatabase) DeletePath(pathRec *record.PathRecord) error {

	deleteRecordSQL := `DELETE from paths WHERE path = ?`
	_, err := d.db.Exec(deleteRecordSQL, pathRec.GetPath())
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *SQLDatabase) GetRecordsByName(optionalPaths []string) ([]*record.PathRecord, error) {
	selectQuery := "SELECT path, score FROM paths WHERE path IN (" + strings.Join((strings.Split(strings.Repeat("?", len(optionalPaths)), "")), ", ") + ")"

	args := make([]interface{}, len(optionalPaths))
	for i, path := range optionalPaths {
		args[i] = path
	}

	rows, err := d.db.Query(selectQuery, args...)
	if err != nil {
		return nil, errors.New("failed to query for paths")
	}

	defer rows.Close()

	var records []*record.PathRecord

	for rows.Next() {
		// var id int
		var path string
		var score int
		// var last_touched any
		err := rows.Scan(&path, &score)
		if err != nil {
			return nil, errors.New("failed to query for paths")
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

		new_rec, _ := record.NewRecord(path, score)
		records = append(records, new_rec)
	}
	return records, nil
}
