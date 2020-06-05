package model

import "github.com/jinzhu/gorm"

type Wanted struct {
	gorm.Model
	LostItem    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Telephony   string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Name        string `gorm:"not null"`
	ImagePath   string `gorm:"not null"`
}
