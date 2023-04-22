package controllers

import (
	"go/todo/initializer"
	"go/todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func PostCreate(c *gin.Context) {
// 	var body struct {
// 		Title       string
// 		Description string
// 		Category    string
// 		Body        string
// 	}

// 	c.Bind(&body)

// 	post := models.Post{
// 		Title:       body.Title,
// 		Describtion: body.Description,
// 		Category:    body.Category,
// 		Body:        body.Body,
// 	}

// 	result := initializer.DB.Create(&post)
// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "successfully created",
// 		"data":    post,
// 	})
// }

func PostCreate(c *gin.Context) {
	// Check Content-Type header
	contentType := c.GetHeader("Content-Type")
	if contentType != "application/json" {
		c.Status(http.StatusBadRequest)
		return
	}

	// Parse request body into Go struct
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Create post in database
	result := initializer.DB.Create(&post)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created",
		"data":    post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializer.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "all datas",
		"data":    posts,
	})
}

func PostShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializer.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "selected data",
		"data":    post,
	})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title       string
		Describtion string
		Category    string
		Body        string
	}
	c.Bind(&body)

	var post models.Post
	initializer.DB.First(&post, id)

	initializer.DB.Model(&post).Updates(models.Post{
		Title:       body.Title,
		Describtion: body.Describtion,
		Category:    body.Category,
		Body:        body.Body,
	})
	c.JSON(200, gin.H{
		"message": "updated data",
		"data":    body,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializer.DB.Delete(models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Sucessfully deleted",
	})
}

func Armagedon(c *gin.Context) {
	// Delete all posts
	initializer.DB.Delete(&[]models.Post{})

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "All posts deleted successfully"})
}
