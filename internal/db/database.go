package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct{
	db 			   *sql.DB
	databaseExists bool
}

func newDatabase(dbFile string) (*Database, error){
	_, err := os.Stat(dbFile)
	databaseExists := !os.IsNotExist(err)

	db ,err := sql.Open("sqlite3", dbFile)
	
	if err != nil{
		return nil, fmt.Errorf("error opening database %v", err)
	}

	return &Database{
		db: db,
		databaseExists: databaseExists,
		}, nil
}

func GetDatabase(dbFile string) (*Database, error) {
	database, err := newDatabase(dbFile)
	if err != nil{
		return nil, fmt.Errorf("error creating or opening database: %v", err)
	}

	if !database.databaseExists {
		err := database.Initizlize()
		if err != nil {
			return nil, fmt.Errorf("error initializing database: %v", err)
		}
	}

	return database, nil

}

func (d *Database) Initizlize() error {
	schemaSQL := `
		CREATE TABLE IF NOT EXISTS paths (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL UNIQUE,
			score INTEGER
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
		INSERT INTO paths (path, score) VALUES (?, ?)
	`
	result, err := d.db.Exec(insertPathSQL, path, 0)

	if err != nil{
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

