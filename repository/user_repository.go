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
