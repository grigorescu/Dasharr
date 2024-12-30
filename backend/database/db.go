package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	// Create a table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS user_stats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tracker_id INTEGER,
		uploaded_torrents INTEGER DEFAULT 0,
		uploaded_amount INTEGER DEFAULT 0,
		snatched INTEGER DEFAULT 0,
		seeding INTEGER DEFAULT 0,
		leeching INTEGER DEFAULT 0,
		ratio REAL DEFAULT 0.0,
		collected_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}
