package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "pustaka-api/book"
    "github.com/go-playground/validator/v10"
    "fmt"
)

type bookHandler struct {
    bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
  return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks (c *gin.Context) {
  books, err := h.bookService.FindAll()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "data" : err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "data" : books,
  })
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
  var bookRequest book.BookRequest

  err := c.ShouldBindJSON(&bookRequest)
  if err != nil {
      errorMessages := []string{}
    for _, e := range err.(validator.ValidationErrors) {
        errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
        errorMessages = append(errorMessages,errorMessage)
    }

     c.JSON(http.StatusBadRequest, gin.H{
         "errors" : errorMessages,
     })
   return
  }

  book, err := h.bookService.Create(bookRequest)

  if err!= nil{
     c.JSON(http.StatusBadRequest, gin.H{
         "errors" : err,
     })
     return
  }

  c.JSON(http.StatusOK, gin.H{
    "data" : book,
  })
}
