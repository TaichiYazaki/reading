package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBNAME = "golang.db"

// Sqlite3設定
func DatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DBNAME), &gorm.Config{})
	return db, err
}
