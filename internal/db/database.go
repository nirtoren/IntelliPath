package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct{
	db *sql.DB
}

func NewDatabase(dbFile string) (*Database, error){
	db ,err := sql.Open("sqlite3", dbFile)
	
	if err != nil{
		return nil, fmt.Errorf("error opening database %v", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Initizlize() error {
	schemaSQL := `
		CREATE TABLE IF NOT EXISTS paths (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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

func (d *Database) InsertPath(path string) (int64, error) {

	insertPathSQL := `
		INSERT INTO paths (path) VALUES (?)
	`
	result, err := d.db.Exec(insertPathSQL, path)

	if err != nil{
		return 0, fmt.Errorf("error on insertion")
	}

	userID, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error while getting last inserted id: %v", err)
	}

	return userID, nil
}