package conf

import (
	"database/sql"
)

// InitializeDB creates the tables if they don't exist
func InitializeDB() {
	db := ConnectDB()
	scenariosDBCreate(db)
	containerDBCreate(db)
	defer db.Close()
}

func scenariosDBCreate(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS scenarios (
		name TEXT PRIMARY KEY,
		state TEXT DEFAULT CREATED
	)`)
	stmt.Exec()

	defer stmt.Close()
}

func containerDBCreate(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS containers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		network TEXT,
		adb_port TEXT,
		serial_port TEXT,
		vnc_port TEXT,
		is_server INTEGER DEFAULT 0,
		state TEXT DEFAULT CREATED
	)`)
	stmt.Exec()

	defer stmt.Close()
}

// ConnectDB returns a connection to the sqlite database
func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}

	return db
}
