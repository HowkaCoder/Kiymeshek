package models

import "gorm.io/gorm"

type Video struct {
  gorm.Model
  ID          uint        `gorm:"primaryKey"`
  Title       string      `gorm:"not null"`
  Description string      `gorm:"not null"`
  Path        string      `gorm:"not null"`
  PostID      uint        `gorm:"default:null"`  
}
