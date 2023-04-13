package db

import "gorm.io/gorm"

type Reading struct {
	gorm.Model
	Title      string
	Price      string
	Review     string
	ImgFile    string
	Impression string
}
