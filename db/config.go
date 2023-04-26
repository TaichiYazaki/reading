package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBNAME = "golang.db"

// Mysql設定
// const (
// 	DBNAME     = "golang"
// 	USERNAME   = "root"
// 	PASSWORD   = ""
// 	IPADDRESS  = "127.0.0.1"
// 	PORTNUMBER = "3306"
// )
// func DatabaseConnection() (*gorm.DB, error) {
// 	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + IPADDRESS + ":" + PORTNUMBER + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	return db, err
// }

// Sqlite3設定
func DatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DBNAME), &gorm.Config{})
	return db, err
}
