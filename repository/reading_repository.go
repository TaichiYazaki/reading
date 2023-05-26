package repository

import (
	"errors"
	"reading/db"
	"reading/model"
)

func ReadingDbInit() error{
	reading := model.Reading{}
	ErrNotFound :=errors.New("ReadingDbInit()でエラー発生")
	db, err := db.DatabaseConnection()
	if err != nil {
		return ErrNotFound
	}
	db.AutoMigrate(&reading)
	return nil
}

func DbCreate(title string, price string, review string, file string, impression string) error{
	reading := model.Reading{Title: title, Price: price, Review: review, ImgFile: file, Impression: impression}
	ErrNotFound := errors.New("DbCreate()でエラー発生")
	db, err := db.DatabaseConnection()
	if err != nil {
		return ErrNotFound
	}
	db.Create(&reading)
	return nil
}

func DbFindAll() ([]model.Reading, error){
	db, err := db.DatabaseConnection()
	ErrNotFound := errors.New("DbCreate()でエラー発生")
	if err != nil {
		return nil, ErrNotFound
	}
	var readings []model.Reading
	db.Find(&readings)
	return readings, nil
}

func DbFindOne(id int) (model.Reading, error) {
	db, err := db.DatabaseConnection()
	ErrNotFound := errors.New("DbFindOne()でエラー発生")
	var reading model.Reading
	if err != nil {
		return reading, ErrNotFound
	}
	db.Find(&reading, id)
	return reading, nil
}

func DbUpdate(id int, title string, price string, review string, impression string) error{
	db, err := db.DatabaseConnection()
	ErrNotFound := errors.New("DbUpdate()でエラー発生")
	if err != nil{
		return ErrNotFound
	}
	var reading model.Reading
	db.First(&reading, id)
	db.Model(&reading).Updates(model.Reading{Title: title, Price: price, Review: review, Impression: impression})
	return nil
}

func DbDelete(id int) error{
	db, err := db.DatabaseConnection()
	ErrNotFound := errors.New("DbDelete()でエラー発生")
	if err != nil{
		return ErrNotFound
	}
	var reading model.Reading
	db.First(&reading, id)
	db.Delete(&reading)
	return nil
}
