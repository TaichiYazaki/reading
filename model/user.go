package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
