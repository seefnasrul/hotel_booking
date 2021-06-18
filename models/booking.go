package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	Email string `gorm:"size:255;not null" json:"email"`
	Name string `gorm:"size:255;not null" json:"name"`
	StartDate time.Time `gorm:"not null" json:"startdate"`
	EndDate time.Time `gorm:"not null" json:"enddate"`
	RoomID uint `gorm:"not null" json:"room_id"` 	
	Room Room
}