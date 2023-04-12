package db

import "gorm.io/gorm"

type Reading struct {
	gorm.Model
	Title      string
	Price      string
	Review     string
	imgFile    string
	impression string
}
