package main

import (
	// "io"
	// "encoding/json"
	"fmt"
	"net/http"
	"book-recom/models"

	controller "book-recom/controllers"

	// "errors"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	// クエリパラメータ取得
	// ex) domain.jp/index?q=xxx
	queryParam := c.Query("q") 
	fmt.Println("=========クエリパラメータ取得", queryParam)
	c.IndentedJSON(http.StatusOK, queryParam)
}



func getCars(c *gin.Context) {

	// inputDataの構造体を用意
	type requestParam struct {
		Text string `json:"text" binding:"required"`
		Id int `json:"id" binding:"required"`
	}

	var json requestParam
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// フォームパラメータ取得
	formParam := c.PostForm("form") 
	fmt.Println("=========フォームパラメータ取得", formParam)
	c.JSON(http.StatusOK, gin.H{"ID": json.Id, "テキスト": json.Text})
}

func getUsers(c *gin.Context) {
	db := models.ConnectDB()

	users := models.GetUsers(db)
	c.JSON(http.StatusOK, gin.H{"users": users})
}


func main() {
	router := gin.Default()

	router.GET("/books", getBooks)

	router.GET("/users", getUsers)

	router.POST("/cars", getCars)


	router.POST("/signup", controller.PostSignUp)
	router.POST("/getSession", controller.GetSession)


	router.Run(":8000")
}