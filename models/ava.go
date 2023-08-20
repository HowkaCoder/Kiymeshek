package models

import "gorm.io/gorm"

type Ava struct {
  gorm.Model
  ID      uint      `gorm:"primaryKey"`
  Path    string    `gorm:"not null"`

}
