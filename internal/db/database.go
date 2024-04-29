package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db	*sql.DB
}

type PathRecord struct {
	Path  string
	Score int8
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

func NewRecord(path string, score int8) (*PathRecord, error) {
	if path == "" {
		return nil, errors.New("path can not be NULL")
	}
	return &PathRecord{Path: path,
		Score: score}, nil
}

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
	schemaSQL := `
		CREATE TABLE IF NOT EXISTS paths (
			path TEXT NOT NULL UNIQUE,
			score INTEGER
		)
	`
	_, err := d.db.Exec(schemaSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}

func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}

	return nil
}

func (d *Database) InsertRecord(pathRec *PathRecord) (int64, error) {

	insertPathSQL := `
		INSERT INTO paths (path, score) VALUES (?, ?)
	`
	result, err := d.db.Exec(insertPathSQL, pathRec.Path, pathRec.Score)

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

func (d *Database) PathSearch(pathToSearch string) (string, int8, error) {

	var path string
	var score int8

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

	return path, score, err
}

func (d *Database) UpdateScore(pathToUpdate string, oldScore int8) error {

	updateScoresSQL := `UPDATE paths SET score = ? WHERE path = ?`
	newScore := oldScore + 1
	_, err := d.db.Exec(updateScoresSQL, newScore, pathToUpdate)
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *Database) DeletePath(pathToDelete string) error {

	deleteRecordSQL := `DELETE from paths WHERE path = ?`
	_, err := d.db.Exec(deleteRecordSQL, pathToDelete)
	if err != nil {
		return errors.New("failed updating the score of a the path")
	}
	return nil
}

func (d *Database) GetRecordsByName(optionalPaths []string) ([]PathRecord, error) {
	selectQuery := "SELECT * FROM paths WHERE path IN (" + strings.Join((strings.Split(strings.Repeat("?", len(optionalPaths)), "")), ", ") + ")"

	args := make([]interface{}, len(optionalPaths))
	for i, path := range optionalPaths {
		args[i] = path
	}

	rows, err := d.db.Query(selectQuery, args...)
	if err != nil {
		return []PathRecord{}, errors.New("failed to query for paths")
	}

	defer rows.Close()

	var records []PathRecord

	for rows.Next() {
		var path string
		var score int8
		err := rows.Scan(&path, &score)
		if err != nil {
			return []PathRecord{}, errors.New("failed to query for paths")
		}
		records = append(records, PathRecord{Path: path, Score: score})

		if err := rows.Err(); err != nil {
			return []PathRecord{}, err
		}
	}
	return records, nil
}
