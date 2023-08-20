package models

import "gorm.io/gorm"

type Category struct {
  gorm.Model 
  ID            uint      `gorm:"primatyKey"`
  Title         string    `gorm:"not null"`
  Description   string    `gorm:"type:varchar(300)"`
  Posts         []Post
}
