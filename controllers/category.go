package controllers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "strconv"
    "go-auth/database"
    "go-auth/models"
)

func GetCategory(c *fiber.Ctx) error {
    categoryID, _ := strconv.Atoi(c.Params("id"))
    var category models.Category

    if err := database.DB.First(&category, categoryID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found", "err": err})
    }

    return c.JSON(category)
}

func CreateCategory(c *fiber.Ctx) error {
    var newCategory models.Category
    if err := c.BodyParser(&newCategory); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    if err := database.DB.Create(&newCategory).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
    }

    return c.Status(fiber.StatusCreated).JSON(newCategory)
}

func IndexCategories(c *fiber.Ctx) error {
    var categories []models.Category
    if err := database.DB.Find(&categories).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found", "error": err.Error()})
    }
    return c.JSON(categories)
}

func UpdateCategory(c *fiber.Ctx) error {
    categoryID, _ := strconv.Atoi(c.Params("id"))
    var category models.Category

    if err := database.DB.First(&category, categoryID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "category not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    var updatedCategory models.Category
    if err := c.BodyParser(&updatedCategory); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    category.Title = updatedCategory.Title
    category.Description = updatedCategory.Description
    if err := database.DB.Save(&category).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
    categoryID, _ := strconv.Atoi(c.Params("id"))
    var category models.Category

    if err := database.DB.First(&category, categoryID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&category).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "category deleted"})
}

