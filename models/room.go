package models

import (
	"github.com/shopspring/decimal"
	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model
	Name  	string `json:"name"`
	Price 	decimal.Decimal `json:"price" sql:"type:decimal(20,2)"`
	HotelID int
}