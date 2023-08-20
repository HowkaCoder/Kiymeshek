package models

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    ID          uint    `gorm:"primaryKey"`
    CategoryID  uint    `gorm:"not null"`
    Title       string  `gorm:"not null"`
    Description string  
    Body1       string  `gorm:"not null"`
    Body2       string
    Body3       string
    Photos      []Photo 
    VideoID     uint    
    Video       Video   `gorm:"foreignKey:VideoID"`
}

