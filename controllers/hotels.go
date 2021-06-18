package controllers

import (
	"net/http"

	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/seefnasrul/go-gin-gorm/models"
)

func GetHotels(c *gin.Context){

	var hotels []models.Hotel

	page := com.StrTo(c.Query("page")).MustInt()
	if page < 1 {
		page = 1
	}


	limit := com.StrTo(c.Query("limit")).MustInt()
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	
	models.DB.Limit(limit).Offset(offset).Order("id desc").Find(&hotels)
	
	var total_data int64
	models.DB.Table("hotels").Count(&total_data)


	c.JSON(http.StatusOK,gin.H{"page":page,"limit":limit,"total":total_data,"data":hotels})

}


type CreateHotelInput struct {
	Name string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func CreateHotel(c *gin.Context){
	var input CreateHotelInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel := models.Hotel{Name: input.Name, Address:input.Address}
	models.DB.Create(&hotel)

	c.JSON(http.StatusOK, gin.H{"data": hotel})

}


func GetHotelByID(c *gin.Context){
	var hotel models.Hotel

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hotel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hotel})
}

//schema validation
type  UpdateHotelInput struct{
	Name  string `json:"name"`
	Address string `json:"address"`  
}

// Update a hotel
func UpdateHotel(c *gin.Context) {
	// Get model if exist
	var hotel models.Hotel
	if err := models.DB.Where("id = ?", c.Param("id")).First(&hotel).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateHotelInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&hotel).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": hotel})
}

// Delete a Hotel
func DeleteHotel(c *gin.Context) {
	// Get model if exist
	var hotel models.Hotel
	if err := models.DB.Where("id = ?", c.Param("id")).First(&hotel).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&hotel)
  
	c.JSON(http.StatusOK, gin.H{"data": true})

}
