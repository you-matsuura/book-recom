package model_redis

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func NewSession(c *gin.Context, cookieKey, redisValue string) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	// 4: 有効期限を示している。0は永久
	if err := conn.Set(c, newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}

	// 3: 有効期限を示している。0はセッションクッキー
	// 4: 有効パスを設定する
	// 5: 有効ドメインを設定する
	// 6,7: セキュリティー系、本番ではtrueにする必要があるかも
	c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

func GetSession(c *gin.Context, cookieKey string) interface{} {
	redisKey, _ := c.Cookie(cookieKey)
	// redisKey := cookieKey
	redisValue, err := conn.Get(c, redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません。")
		return nil
	case err != nil:
		fmt.Println("Session取得時にエラー発生：" + err.Error())
		return nil
	}
	return redisValue
}

func DeleteSession(c *gin.Context, cookieKey string) {
	redisId, _ := c.Cookie(cookieKey)
	conn.Del(c, redisId)
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}
