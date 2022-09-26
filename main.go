package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:tmDevFlats0987^(@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	router := gin.Default()

	if err != nil {
		log.Fatal("DB Connection Error")
		log.Fatal("Oke")
	}

	db.AutoMigrate(&book.Book{})

  bookRepository := book.NewRepository(db)

	// book := book.Book{}
	// book.Title = "Manusia Harimau"
	// book.Price = 50000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Cerita yang tidak masuk akal, namun diterima oleh rakyat Indonesia"

	// err = db.Create(&book).Error
	// if err != nil {
  // fmt.Println("====================")
  // fmt.Println("Error Createing Book")
	// fmt.Println("====================")
	// }

  // var book book.Book

  // ==== Get ====
  // err = db.Where("id = ? ", 1).First(&book).Error
  // if err != nil {
      // fmt.Println("====================")
      // fmt.Println("Error Get Book")
      // fmt.Println("====================")
  // }

  // ==== Delete ====
  // err = db.Delete(&book).Error
  // if err != nil {
    // fmt.Println("====================")
    // fmt.Println("Error Deleting Book")
    // fmt.Println("====================")
  }

  // ==== Update ====
  // book.Title = "Man Tiger (2)"
  // err = db.Save(&book).Error
  // if err != nil {
      // fmt.Println("====================")
      // fmt.Println("Error Updating Book")
      // fmt.Println("====================")
  // }

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)

	router.Run(":8889")
}

//  01:50:30 - 02:01:00 Insterface
