package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name     string `gorm:"not null" form:"name" binding:"required"`
	Email    string `gorm:"not null" form:"email" binding:"required"`
	Password string `gorm:"not null" form:"password" binding:"required"`
}
