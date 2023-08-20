package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)


type User struct {
  gorm.Model
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"not null" json:"name"`
    Email    string `gorm:"not null;unique" json:"email"`
    Password string `gorm:"not null" json:"password"`
    AvaID    *uint   `json:"ava_id"` // Use "gorm" tag for foreign key constraint
    Ava      Ava    `gorm:"foreignKey:AvaID"`
    Role     string `gorm:"not null"`
}


type JWTClaims struct {
  ID      uint      `json:"id"`
  Role    string    `json:"role"`
  jwt.StandardClaims
}
