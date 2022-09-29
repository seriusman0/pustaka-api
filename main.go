package main

import (
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
  bookService := book.NewService(bookRepository)
  bookHandler := handler.NewBookHandler(bookService)

  bookRequest := book.BookRequest{
    Title: "Gundam",
    Price: "200000",
  }

  bookService.Create(bookRequest)

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run(":8888")
}

//  02:47:02 - --:--:-- Perbaiki service.go untuk memetakan book

