package main

import (
	"go-auth/database"
	"go-auth/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
  app := fiber.New()
  app.Use(logger.New())
	app.Use(cors.New())

  database.ConnectToDB()
  routes.SetupRoutes(app)
  if err := app.Listen("0.0.0.0:8080"); err != nil {
    log.Fatal(err)
  }

}
