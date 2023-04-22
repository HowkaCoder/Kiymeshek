package migrate

import (
	"go/todo/initializer"
	"go/todo/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func Migrate() {
	initializer.DB.AutoMigrate(&models.Post{})
}
