package routes

import (
	"go-auth/controllers"
  "go-auth/database"
  "go-auth/config"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
  app.Get("/hello",func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "message" : "Hello world",
    })
  })
  app.Get("/boom", func(c *fiber.Ctx) error {
    io := "DROP DATABASE IF EXISTS " + config.Config("DB_NAME")

		res := database.DB.Exec(io)
		if res.Error != nil {
			return res.Error
		}
    io = "CREATE DATABASE "+config.Config("DB_NAME")
    res = database.DB.Exec(io)
    if res.Error != nil {
      return res.Error
    }
    database.ConnectToDB()

		return c.SendString("Database reborned")
	})
  


  app.Post("/register", controllers.Register)
  app.Post("/login" ,controllers.Login)


  app.Get("/users/:id" , controllers.Auth("admin"),controllers.GetUser)
  app.Put("/users/:id" ,controllers.Auth("user"), controllers.UpdateUser)
  app.Delete("/users/:id" ,controllers.Auth("admin"), controllers.DeleteUser)

  app.Get("/videos/:id" , controllers.GetVideo)
  app.Get("/videos" , controllers.IndexVideo)
  app.Post("/videos" , controllers.Auth("admin"),controllers.CreateVideo)
  app.Put("/videos/:id" , controllers.Auth("admin") , controllers.UpdateVideo)
  app.Delete("/videos/:id" , controllers.Auth("admin") , controllers.DeleteVideo)

  app.Get("/photos" , controllers.IndexPhoto)
  app.Post("/photos" , controllers.Auth("admin") , controllers.CreatePhoto)
  app.Get("/photos/:id"   , controllers.GetPhoto)
  app.Put("/photos/:id" , controllers.Auth("admin") , controllers.UpdatePhoto)
  app.Delete("/photos/:id" , controllers.Auth("admin") , controllers.DeletePhoto)


  app.Get("/categories" , controllers.IndexCategories)
  app.Get("/categories/:id" , controllers.GetCategory)
  app.Post("/categories" , controllers.Auth("admin") , controllers.CreateCategory)
  app.Put("/categories/:id" , controllers.Auth("admin") ,controllers.UpdateCategory)
  app.Delete("/categories/:id" , controllers.Auth("admin") , controllers.DeleteCategory)



  app.Get("/posts", controllers.IndexPosts)
  app.Get("/posts/:id", controllers.GetPost)
  app.Post("/posts", controllers.Auth("admin"),controllers.CreatePost)
  app.Put("/posts/:id", controllers.Auth("admin"),controllers.UpdatePost)
  app.Delete("/posts/:id", controllers.Auth("admin"),controllers.DeletePost)


  app.Post("/avas" ,controllers.Auth("user"), controllers.CreateAva)
  app.Get("/avas/:id"  ,controllers.Auth("user"),controllers.GetAva)
  app.Put("/avas/:id" ,controllers.Auth("user"), controllers.UpdateAva)
  app.Delete("/avas/:id" ,controllers.Auth("user"), controllers.DeleteAva)

  

}
