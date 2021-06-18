package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seefnasrul/go-gin-gorm/models"
)

func FindBooks(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data":books})
}



//schema validation
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
	
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": book})

}

func FindBook(c *gin.Context) {  // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//schema validation
type  UpdateBookInput struct{
	Title  string `json:"title"`
	Author string `json:"author"`  
}

// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&book).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
}
