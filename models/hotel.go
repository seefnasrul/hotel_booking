package models

import (
	"github.com/jinzhu/gorm"
)

type Hotel struct {
	gorm.Model
	Name  	string `json:"name"`
	Address string `json:"address"`
	Rooms []Room
}