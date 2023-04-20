package db

import (
	"log"
)

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
