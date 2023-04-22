package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `gorm:"type:text"`
	Describtion string `gorm:"type:text"`
	Category    string `gorm:"type:text"`
	Body        string `gorm:"type:text"`
}
