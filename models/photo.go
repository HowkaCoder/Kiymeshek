package models

import "gorm.io/gorm"


type Photo struct {
  gorm.Model
  ID            uint      `gorm:"primaryKey"`
  Title         string    `gorm:"not null"`
  Description   string    `gorm:"type:varchar(300)"`
  Path          string    `gorm:"not null"`
  PostID        uint      `gorm:"default:null"` 
}
