package db

import (
	"log"

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

func DbInit() {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Reading{})
}

func DbCreate(title string, price string, review string, file string, impression string) {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&Reading{Title: title, Price: price, Review: review, ImgFile: file, Impression: impression})
}

func DbFindAll() []Reading {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var readings []Reading
	db.Find(&readings)
	return readings
}

func DbFindOne(id int) Reading {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading Reading
	db.Find(&reading, id)
	return reading
}

func DbUpdate(id int, title string, price string, review string, impression string) {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading Reading
	db.First(&reading, id)
	db.Model(&reading).Updates(Reading{Title: title, Price: price, Review: review, Impression: impression})
}

func DbDelete(id int) {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading Reading
	db.First(&reading, id)
	db.Delete(&reading)

}
