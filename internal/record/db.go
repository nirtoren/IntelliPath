package record

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db	*sql.DB
}


func NewDatabase(dbFile string) (*Database, error) {
	_, err := os.Stat(dbFile)
	databaseExists := !os.IsNotExist(err)

	if databaseExists {
		fmt.Println("Seems like you aleadt have a database, try to remove it and re-run.")
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, errors.New("an error occurred while creating the database")
	}

	return &Database{
		db:	db,
	}, nil
}

// func NewRecord(path string, score int8) (*PathRecord, error) {
// 	if path == "" {
// 		return nil, errors.New("path can not be NULL")
// 	}
// 	return &PathRecord{Path: path,
// 		Score: score}, nil
// }

func GetDatabase(dbFile string) (*Database, error) {
	if info, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Println("A database was not inititalized. Run 'intellipath init'")
		fmt.Println(info)
		os.Exit(1)
		return &Database{
			db:	nil,
		}, errors.New("A database was not inititalized. Run 'intellipath init'")
	} else {
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			fmt.Println("Error: Could not get a database.")
			os.Exit(1)
		}

		return &Database{
			db:	db,
		}, nil
	}
}

func (d *Database) Initizlize() error {
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

func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}

	return nil
}

func ParallelCleanUp(d *Database, resultCh chan<-error){
	cutOffTime := time.Now().Add(-5 * 24 * time.Hour)

	_, err := d.db.Exec("DELETE FROM paths WHERE last_touched < ?", cutOffTime)
	resultCh <- err
}

func (d *Database) InsertRecord(pathRec *PathRecord) (int64, error) {

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

func (d *Database) GetAllPaths() ([]string, error) {

	var paths []string

	getAllPathsSQL := `
		SELECT path FROM paths
	`
	rows, err := d.db.Query(getAllPathsSQL)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, fmt.Errorf("error scanning row %v", err)
		}
		paths = append(paths, path)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	return paths, nil
}

func (d *Database) PathSearch(pathToSearch string) (*PathRecord, error) {

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

	newRec,_ := NewRecord(path, score)

	return newRec, err
}

func (d *Database) UpdateScore(pathRec *PathRecord) error {

	updateScoresSQL := `UPDATE paths SET score = ? WHERE path = ?`
	newScore := pathRec.GetScore() + 1
	_, err := d.db.Exec(updateScoresSQL, newScore, pathRec.GetPath())
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *Database) DeletePath(pathRec *PathRecord) error {

	deleteRecordSQL := `DELETE from paths WHERE path = ?`
	_, err := d.db.Exec(deleteRecordSQL, pathRec.GetPath())
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *Database) GetRecordsByName(optionalPaths []string) ([]*PathRecord, error) {
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

	var records []*PathRecord

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

		new_rec, _ := NewRecord(path, score)
		records = append(records, new_rec)
	}
	return records, nil
}
