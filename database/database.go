package database 
 
import ( 
 "fmt" 
 "go-auth/models" 
 "log"  
 "gorm.io/driver/postgres" 
 "gorm.io/gorm" 
) 
 
var DB *gorm.DB 
 
func ConnectToDB() { 
 var err error 
 dsn := "host=192.168.88.126 user=howka password=euF@I16TuradPUXF dbname=kiymeshek port=3306 sslmode=disable TimeZone=Asia/Shanghai"

 DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
 if err != nil { 
  log.Fatal(err) 
 } 
 DB.AutoMigrate(&models.User{}) 
 DB.AutoMigrate(&models.Ava{}) 
 DB.AutoMigrate(&models.Photo{}) 
 DB.AutoMigrate(&models.Video{}) 
 DB.AutoMigrate(&models.Category{}) 
 DB.AutoMigrate(&models.Post{}) 
 fmt.Println("Database Migrated") 
 
}
