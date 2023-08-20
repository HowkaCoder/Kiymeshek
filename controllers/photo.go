package controllers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "strconv"
    "go-auth/database"
    "go-auth/models"
  )

// ...

func GetPhoto(c *fiber.Ctx) error {
    photoID , _ := strconv.Atoi(c.Params("id"))
    var photo models.Photo

    if err := database.DB.First(&photo, photoID).Error; err != nil {
    
      return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Photo not found" , "err":err})
  }

    return c.JSON(photo)
}

func CreatePhoto(c *fiber.Ctx) error {
    photoFile, err := c.FormFile("photo")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    photoPath := "uploads/photo/" + photoFile.Filename
    if err := c.SaveFile(photoFile, photoPath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo"})
    }

    var newPhoto models.Photo
    if err := c.BodyParser(&newPhoto); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"invalid form data"})
		}
    newPhoto.Path = photoPath
    if err := database.DB.Create(&newPhoto).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error 1"})
    }


    return c.Status(fiber.StatusCreated).JSON(newPhoto)
}

func IndexPhoto(c *fiber.Ctx) error {
    var photos []models.Photo
    if err := database.DB.Find(&photos).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found", "error": err.Error()})
    }
    return c.JSON(photos)
}



func UpdatePhoto(c *fiber.Ctx) error {
    photoID , _ := strconv.Atoi(c.Params("id"))
    var photo models.Photo

    if err := database.DB.First(&photo, photoID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "photo not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }
    
    photoFile, err := c.FormFile("photo")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    photoPath := "uploads/photo/" + photoFile.Filename
    if err := c.SaveFile(photoFile, photoPath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo"})
    }



  

    var updatedPhoto models.Photo
    if err := c.BodyParser(&updatedPhoto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    photo.Path = photoPath
    photo.Title = updatedPhoto.Title
    photo.Description = updatedPhoto.Description

    if err := database.DB.Save(&photo).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(photo)
}

func DeletePhoto(c *fiber.Ctx) error {
    photoID  , _:= strconv.Atoi(c.Params("id"))
    var photo models.Photo

    if err := database.DB.First(&photo, photoID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Photo not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&photo).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "photo deleted"})
}

