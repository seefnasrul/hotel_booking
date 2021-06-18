package main

import (
	"reflect"
	"sync"
  	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/seefnasrul/go-gin-gorm/models"
	"github.com/seefnasrul/go-gin-gorm/controllers"
	"github.com/seefnasrul/go-gin-gorm/middlewares"
	"github.com/seefnasrul/go-gin-gorm/utils"
	"github.com/go-playground/validator/v10"
)

func main() {

	binding.Validator = new(defaultValidator)

	//register custom validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", utils.BookableDate)
	}

  	r := gin.Default()

	models.ConnectDataBase()

	public := r.Group("/api")
	admin := r.Group("/api/admin")
	admin.Use(middlewares.SetMiddlewareAuthentication())

	public.POST("/login",controllers.Login)
	public.POST("/register",controllers.Register)
	public.GET("/search-room",controllers.SearchAvailableRoom)


	admin.GET("/books",controllers.FindBooks)
	admin.POST("/books", controllers.CreateBook)
	admin.GET("/books/:id", controllers.FindBook)
	admin.PATCH("/books/:id", controllers.UpdateBook)
	admin.DELETE("/books/:id", controllers.DeleteBook)


	//hotels CRUD
	admin.GET("/hotels",controllers.GetHotels)
	admin.POST("/hotels",controllers.CreateHotel)
	admin.GET("/hotels/:id", controllers.GetHotelByID)
	admin.PATCH("/hotels/:id", controllers.UpdateHotel)
	admin.DELETE("/hotels/:id", controllers.DeleteHotel)


	//rooms CRUD
	admin.GET("/hotels/:id/rooms",controllers.GetRooms)
	admin.POST("/hotels/:id/rooms",controllers.CreateRoom)
	admin.GET("/hotels/:id/rooms/:roomid",controllers.GetRoomByID)
	admin.PATCH("/hotels/:id/rooms/:roomid", controllers.UpdateRoom)
	admin.DELETE("/hotels/:id/rooms/:roomid", controllers.DeleteRoom)

	r.Run(":8070")
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}