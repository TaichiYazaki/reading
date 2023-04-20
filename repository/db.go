package repository

import (
	"log"
	"reading/db"
	"reading/model"
)

func DbInit() {
	reading := model.Reading{}
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&reading)
}

func DbCreate(title string, price string, review string, file string, impression string) {
	reading := model.Reading{Title: title, Price: price, Review: review, ImgFile: file, Impression: impression}
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&reading)
}

func DbFindAll() []model.Reading {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var readings []model.Reading
	db.Find(&readings)
	return readings
}

func DbFindOne(id int) model.Reading {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading model.Reading
	db.Find(&reading, id)
	return reading
}

func DbUpdate(id int, title string, price string, review string, impression string) {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading model.Reading
	db.First(&reading, id)
	db.Model(&reading).Updates(model.Reading{Title: title, Price: price, Review: review, Impression: impression})
}

func DbDelete(id int) {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	var reading model.Reading
	db.First(&reading, id)
	db.Delete(&reading)

}
