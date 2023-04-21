package repository

import (
	"log"
	"reading/db"
	"reading/model"
)

func UserDbInit() {
	user := model.User{}
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&user)
}

func UserCreate(name string, email string, password string) {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{Name: name, Email: email, Password: password}
	db.Create(&user)
}
