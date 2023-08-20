package controllers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "strconv"
    "go-auth/database"
    "go-auth/models"
)

func GetPost(c *fiber.Ctx) error {
    postID, _ := strconv.Atoi(c.Params("id"))
    var post models.Post

    if err := database.DB.Preload("Photos").Preload("Video").First(&post, postID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found", "err": err})
    }

    return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
    var newPost models.Post
    if err := c.BodyParser(&newPost); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    if err := database.DB.Create(&newPost).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
    }

    return c.Status(fiber.StatusCreated).JSON(newPost)
}

func IndexPosts(c *fiber.Ctx) error {
    var posts []models.Post
    if err := database.DB.Find(&posts).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found", "error": err.Error()})
    }
    return c.JSON(posts)
}

func UpdatePost(c *fiber.Ctx) error {
    postID, _ := strconv.Atoi(c.Params("id"))
    var post models.Post

    if err := database.DB.First(&post, postID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "post not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    var updatedPost models.Post
    if err := c.BodyParser(&updatedPost); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    post.CategoryID = updatedPost.CategoryID
    post.Title = updatedPost.Title
    post.Description = updatedPost.Description
    post.Body1 = updatedPost.Body1
    post.Body2 = updatedPost.Body2
    post.Body3 = updatedPost.Body3
    post.VideoID = updatedPost.VideoID
    if err := database.DB.Save(&post).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
    postID, _ := strconv.Atoi(c.Params("id"))
    var post models.Post

    if err := database.DB.First(&post, postID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&post).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "post deleted"})
}

