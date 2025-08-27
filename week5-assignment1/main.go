package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Price  int    `json:"price"`
    Stock  int    `json:"stock"`
}

var books = []Book{
    {ID: "1", Title: "จดหมายจากดาวแมว", Author: "นทธี ศศิวิมล", Price: 220, Stock: 10},
    {ID: "2", Title: "กำเนิดบาปทาโกปี้", Author: "TAIZAN5", Price: 149, Stock: 10},
    {ID: "3", Title: "EMOTIONS 101 วิชาจัดการอารมณ์", Author: "อินฮยอนจิน", Price: 216, Stock: 15},
}

func getBooks(c *gin.Context) {
	IDQuery := c.Query("ID")
	if IDQuery != "" {
		filter := []Book{}
		for _, book := range books {
			if fmt.Sprint(book.ID) == IDQuery {
				filter = append(filter, book)
			}
		}

		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, books)
}
func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"massage": "healthy"})
		fmt.Println()
	})
	api := r.Group("/api/v1")
	{
		api.GET("/bookShop", getBooks)
		fmt.Println()
	}

	r.Run(":8080")
}

