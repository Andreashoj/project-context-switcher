package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./project-context-switcher.db")
	if err != nil {
		return nil, fmt.Errorf("opening the database failed: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging the database failed: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	if err = runMigrations(db); err != nil {
		return nil, fmt.Errorf("running the migrations failed: %w", err)
	}

	fmt.Println("connected to the database")

	return db, nil
}

func runMigrations(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS projects (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		name VARCHAR(255) NOT NULL,
    		path VARCHAR(255) NOT NULL,
		    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TRIGGER IF NOT EXISTS projects_updated_at AFTER UPDATE ON projects 
		BEGIN 
			UPDATE projects SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
		END`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("something went wrong with one of the migrations: %w", err)
		}
	}

	return nil
}
