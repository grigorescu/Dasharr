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
    downloaded_amount INTEGER DEFAULT 0,
    snatched INTEGER DEFAULT 0,
    seeding INTEGER DEFAULT 0,
    leeching INTEGER DEFAULT 0,
    ratio REAL DEFAULT 0.0,
    required_ratio REAL DEFAULT 0.0,
    last_access DATETIME DEFAULT NULL,
    torrent_comments INTEGER DEFAULT 0,
    invited INTEGER DEFAULT 0,
    forum_posts INTEGER DEFAULT 0,
    warned BOOLEAN DEFAULT 0,
    class TEXT DEFAULT NULL,
    donor BOOLEAN DEFAULT 0,
    uploaded_rank INTEGER DEFAULT 0,
    downloaded_rank INTEGER DEFAULT 0,
    uploads_rank INTEGER DEFAULT 0,
    requests_rank INTEGER DEFAULT 0,
    bounty_rank INTEGER DEFAULT 0,
    posts_rank INTEGER DEFAULT 0,
    artists_rank INTEGER DEFAULT 0,
    overall_rank INTEGER DEFAULT 0,
    collected_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}
