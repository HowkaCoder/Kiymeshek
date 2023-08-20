package controllers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "strconv"
    "go-auth/database"
    "go-auth/models"
    "fmt"
  )

// ...

func GetAva(c *fiber.Ctx) error {
    avaID , _ := strconv.Atoi(c.Params("id"))
    var ava models.Ava

    if err := database.DB.First(&ava, avaID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ava not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(ava)
}

func CreateAva(c *fiber.Ctx) error {
    avatarFile, err := c.FormFile("avatar")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    avatarPath := "uploads/" + avatarFile.Filename
    if err := c.SaveFile(avatarFile, avatarPath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save avatar"})
    }

    var newAva models.Ava
    newAva.Path = avatarPath

    if err := database.DB.Create(&newAva).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error 1"})
    }

    userID, err := strconv.Atoi(c.FormValue("user_id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid user ID"})
    }

    var user models.User
    if err := database.DB.Find(&user, userID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
    }

     user.AvaID = &newAva.ID

    if err := database.DB.Save(&user).Error; err != nil {
        fmt.Println("Error saving user:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
    }


    return c.Status(fiber.StatusCreated).JSON(newAva)
}


func UpdateAva(c *fiber.Ctx) error {
    avaID , _ := strconv.Atoi(c.Params("id"))
    var ava models.Ava

    if err := database.DB.First(&ava, avaID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ava not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    var updatedAva models.Ava
    if err := c.BodyParser(&updatedAva); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    ava.Path = updatedAva.Path
    if err := database.DB.Save(&ava).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(ava)
}

func DeleteAva(c *fiber.Ctx) error {
    avaID  , _:= strconv.Atoi(c.Params("id"))
    var ava models.Ava

    if err := database.DB.First(&ava, avaID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ava not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&ava).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "Ava deleted"})
}
