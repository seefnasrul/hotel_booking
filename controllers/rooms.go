package controllers

import (
	"net/http"
	"strconv"
	"errors"
	"fmt"
	// "log"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/seefnasrul/go-gin-gorm/models"
	"github.com/seefnasrul/go-gin-gorm/utils"
)



// type ValidationError struct {
// 	Field  string `json:"field"`
// 	Reason string `json:"reason"`
// }

// func Descriptive(verr validator.ValidationErrors) []ValidationError {
// 	errs := []ValidationError{}

// 	for _, f := range verr {
// 		err := f.ActualTag()
// 		if f.Param() != "" {
// 			err = fmt.Sprintf("%s=%s", err, f.Param())
// 		}
// 		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
// 	}

// 	return errs
// }

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}



func GetRooms(c *gin.Context){

	var rooms []models.Room
	var hotel models.Hotel

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hotel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hotel ID!"})
		return
	}

	page := com.StrTo(c.Query("page")).MustInt()
	if page < 1 {
		page = 1
	}


	limit := com.StrTo(c.Query("limit")).MustInt()
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	
	//convert param to int
	hotel_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//select raw
	models.DB.Raw("SELECT id, name, price, created_at, updated_at, hotel_id FROM rooms WHERE hotel_id = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT ? OFFSET ?", hotel_id,limit,offset).Scan(&rooms)

	var total_data int64
	//count to get total of data
	models.DB.Table("rooms").Where(&models.Room{HotelID:hotel_id}).Count(&total_data)

	c.JSON(http.StatusOK,gin.H{"page":page,"limit":limit,"total":total_data,"data":rooms})

}


type CreateRoomInput struct {
	Name string `json:"name" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required"`
}

func CreateRoom(c *gin.Context){

	var input CreateRoomInput
	var room models.Room
	var hotel models.Hotel

	if err := models.DB.Where("id = ?", c.Param("id")).First(&hotel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	room = models.Room{Name: input.Name, Price:input.Price, HotelID:hotel_id}
	models.DB.Create(&room)

	c.JSON(http.StatusOK, gin.H{"data": room})

}


func GetRoomByID(c *gin.Context){
	var room models.Room

	if err := models.DB.Where("id = ?", c.Param("roomid")).Where("hotel_id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

// //schema validation
type  UpdateRoomInput struct{
	Name  string `json:"name"`
	Price decimal.Decimal `json:"address"`  
	HotelID int `json:"hotel_id"`
}

// Update a hotel
func UpdateRoom(c *gin.Context) {
	// Get model if exist
	var room models.Room
	var hotel models.Hotel

	if err := models.DB.Where("id = ?", c.Param("roomid")).Where("hotel_id = ?", c.Param("id")).First(&room).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateRoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.HotelID > 0{
		if err := models.DB.Where("id = ?",input.HotelID).First(&hotel).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hotel ID!"})
			return
		}
	}
	

	models.DB.Model(&room).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": room})
}

// Delete a Hotel
func DeleteRoom(c *gin.Context) {
	// Get model if exist
	var room models.Room
	if err := models.DB.Where("id = ?", c.Param("roomid")).Where("hotel_id = ?", c.Param("id")).First(&room).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&room)
  
	c.JSON(http.StatusOK, gin.H{"data": true})

}
