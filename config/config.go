package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string{
  if err := godotenv.Load(".env"); err != nil {                               
   log.Fatal(err)                                                                  }              
  return os.Getenv(key)
}
