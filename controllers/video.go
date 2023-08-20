package controllers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "strconv"
    "go-auth/database"
    "go-auth/models"
)

func GetVideo(c *fiber.Ctx) error {
    videoID, _ := strconv.Atoi(c.Params("id"))
    var video models.Video

    if err := database.DB.First(&video, videoID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Video not found", "err": err})
    }

    return c.JSON(video)
}

func CreateVideo(c *fiber.Ctx) error {
    videoFile, err := c.FormFile("video")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    videoPath := "uploads/video/" + videoFile.Filename
    if err := c.SaveFile(videoFile, videoPath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save video"})
    }

    var newVideo models.Video
    if err := c.BodyParser(&newVideo); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }
    newVideo.Path = videoPath
    if err := database.DB.Create(&newVideo).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
    }

    return c.Status(fiber.StatusCreated).JSON(newVideo)
}

func IndexVideo(c *fiber.Ctx) error {
    var videos []models.Video
    if err := database.DB.Find(&videos).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found", "error": err.Error()})
    }
    return c.JSON(videos)
}

func UpdateVideo(c *fiber.Ctx) error {
    videoID, _ := strconv.Atoi(c.Params("id"))
    var video models.Video

    if err := database.DB.First(&video, videoID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "video not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    videoFile, err := c.FormFile("video")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
    }

    videoPath := "uploads/video/" + videoFile.Filename
    if err := c.SaveFile(videoFile, videoPath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save video"})
    }

    var updatedVideo models.Video
    if err := c.BodyParser(&updatedVideo); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    video.Path = videoPath
    video.Title = updatedVideo.Title
    video.Description = updatedVideo.Description
    if err := database.DB.Save(&video).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(video)
}

func DeleteVideo(c *fiber.Ctx) error {
    videoID, _ := strconv.Atoi(c.Params("id"))
    var video models.Video

    if err := database.DB.First(&video, videoID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Video not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    if err := database.DB.Delete(&video).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"message": "video deleted"})
}

