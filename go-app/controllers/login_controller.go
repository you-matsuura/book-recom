package controller

import (
	// model_db "go_blog/model/db"
	model_redis "book-recom/models/redis"
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
)

// 新規会員登録
func PostSignUp(c *gin.Context) {
	// id := c.PostForm("user_id")
	// pw := c.PostForm("password")
	// user, err := model_db.Signup(c, id, pw)
	// if err != nil {
	// 	c.Redirect(http.StatusFound, "/signup")
	// 	return
	// }
	cookieKey := "session_id"
	model_redis.NewSession(c, cookieKey, "セッションバリュー")
	c.JSON(http.StatusOK, gin.H{"message": "セッションに保存完了"})
	// c.Redirect(http.StatusFound, "/")
}

func GetSession(c *gin.Context) {
	cookieKey := "session_id"
	redisValue := model_redis.GetSession(c, cookieKey)

	c.JSON(http.StatusOK, gin.H{"redisValue": redisValue})
}