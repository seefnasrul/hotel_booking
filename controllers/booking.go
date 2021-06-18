package controllers

import (
	"net/http"
	// "strconv"
	// "encoding/json"
	"errors"
	"time"
	// "fmt"
	// "log"

	"github.com/go-playground/validator/v10"
	// "github.com/shopspring/decimal"
	// "github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/seefnasrul/go-gin-gorm/models"
	"github.com/seefnasrul/go-gin-gorm/utils"
)



type SearchAvailableRoomInput struct {
	StartDate time.Time `form:"startdate" binding:"required,bookabledate" time_format:"2006-01-02" time_utc:"1"`
	EndDate time.Time `form:"enddate" binding:"required,gtfield=StartDate" time_format:"2006-01-02" time_utc:"1"`
}

func SearchAvailableRoom(c *gin.Context){
	
	var input SearchAvailableRoomInput

	if err := c.ShouldBind(&input); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var rooms []models.Room
	models.DB.Raw("SELECT id, name, price, hotel_id FROM rooms WHERE deleted_at IS NULL AND id NOT IN (SELECT room_id FROM bookings WHERE start_date <= ? AND end_date > ? OR start_date < ? AND end_date >= ?)",input.StartDate,input.StartDate,input.EndDate,input.EndDate).Scan(&rooms)
	c.JSON(http.StatusOK, gin.H{"data": rooms})

}
