package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "pustaka-api/book"
    "github.com/go-playground/validator/v10"
    "fmt"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Seriusman Waruwu",
		"bio":  "A Software engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello",
		"subtitle": "Saya mau belajar Golang",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id": id, "title" : title,
	})
}

func QueryHandler(c *gin.Context) {
	price := c.Query("price")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBookHandler(c *gin.Context) {
  var bookInput book.BookInput

  err := c.ShouldBindJSON(&bookInput)
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

  c.JSON(http.StatusOK, gin.H{
    "title" : bookInput.Title,
    "price" : bookInput.Price,
  })
}
