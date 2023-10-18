package main

import (
	"go-auth/database"
	"go-auth/routes"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
func main(){
  app := fiber.New()
  app.Use(logger.New())
	app.Use(cors.New())

  database.ConnectToDB()
  routes.SetupRoutes(app)
  if err := app.Listen(getPort()); err != nil {
    log.Fatal(err)
  }

}
