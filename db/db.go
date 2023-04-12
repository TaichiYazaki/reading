package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBNAME     = "study"
	USERNAME   = "root"
	PASSWORD   = ""
	IPADDRESS  = "127.0.0.1"
	PORTNUMBER = "3306"
)

func DatabaseConnection() (*gorm.DB, error) {
	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + IPADDRESS + ":" + PORTNUMBER + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
