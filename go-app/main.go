package main

import (
	// "io"
	// "encoding/json"
	"book-recom/models"
	"errors"
	"fmt"
	"net/http"

	controller "book-recom/controllers"

	// "errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		Id   int    `json:"id" binding:"required"`
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
	type User struct {
		Id    int    `gorm:"column:msu_user_id"`
		Email string `gorm:"column:msu_email"`
	}

	db := models.NewSQLHandler()
	users := []User{}

	rows := db.DB.Table("mst_user").Select("msu_user_id", "msu_email").Find(&users)

	// check error ErrRecordNotFound
	if errors.Is(rows.Error, gorm.ErrRecordNotFound) {
		fmt.Println("データベース接続失敗")
		// log.Fatal(rows.Error)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func main() {

	models.DBOpen()
	defer models.DBClose()

	router := gin.Default()

	router.GET("/books", getBooks)

	router.GET("/users", getUsers)

	router.POST("/cars", getCars)

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	// router.POST("/getSession", controller.GetSession)

	router.Run(":8000")
}
