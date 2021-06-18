package utils

import (
	"time"
	"github.com/go-playground/validator/v10"
	"fmt"
)

var BookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	fmt.Println(date,ok)
	if ok {
		
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}