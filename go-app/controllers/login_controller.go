package controller

import (
	// model_db "go_blog/model/db"
	crypto "book-recom/crypto"
	redis "book-recom/models/redis"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"book-recom/models"

	// "os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"column:msu_user_id"`
	NickName  string    `gorm:"column:msu_nick_name"`
	Email     string    `gorm:"column:msu_email"`
	Password  string    `gorm:"column:msu_hash_password"`
	CreatedAt time.Time `gorm:"column:msu_created_at"`
	UpdatedAt time.Time `gorm:"column:msu_updated_at"`
}

type ParamSignUp struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// メールアドレスとパワスワードのバリデーションチェック
func (p ParamSignUp) invalid() bool {

	fmt.Println("==============パラメータ", p.Email, p.Password)

	// メールアドレスのバリデーション
	regexPattern := `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`
	r := regexp.MustCompile(regexPattern)
	if ok := r.MatchString(p.Email); ok == false {
		fmt.Println(p.Email)
		fmt.Println("=================メアドが正しくないです", p.Email)
		return true
	}

	// パスワードのバリデーション
	const minPassLength int = 6 // パスワードの最低文字数
	if pwLength := len(p.Password); pwLength < minPassLength {
		fmt.Println("=================パスワードが不適切です", pwLength)
		return true
	}

	return false
}

// 新規会員登録
// - 新規ユーザーの作成
func SignUp(c *gin.Context) {
	var paramSignUp ParamSignUp

	// jsonからデータを受け取る
	if err := c.ShouldBindJSON(&paramSignUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "jsonからの変換に失敗しました"})
		return
	}

	// メールアドレスとパワスワードのバリデーションチェック
	if paramSignUp.invalid() == true {
		c.JSON(http.StatusBadRequest, gin.H{"message": "メールアドレスまたはパスワードが不適切です"})
		return
	}

	// パラメータのチェック完了
	db := models.NewSQLHandler()

	var checkDupEmail User
	user := User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println("============0", paramSignUp)
	// メールアドレスが既に存在するとき、処理を抜ける

	if err := db.DB.Table("mst_user").Where("msu_email = ?", paramSignUp.Email).First(&checkDupEmail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 新規メールアドレスの場合は、新しくデータを作成する
			user.Email = paramSignUp.Email
		} else {
			// データベースエラー
			fmt.Println("=============会員登録時にエラーが発生しました", err)
			return
		}
	} else {
		// ユーザーが既に存在する場合はエラーメッセージを返す
		fmt.Println("=============既にそのメールアドレスは保存されています")
		return
	}

	// パスワードをハッシュ化する
	if hashPassword, err := crypto.PasswordEncrypt(paramSignUp.Password); err != nil {
		fmt.Println("パスワードのハッシュ化に失敗しました", err)
		return
	} else {
		user.Password = hashPassword
	}

	fmt.Println("============保存直前のデータ", user)
	result := db.DB.Table("mst_user").Create(&user)
	if result.Error != nil {
		panic("=============会員登録に失敗しました")
	}
	c.JSON(http.StatusOK, gin.H{"message": "会員登録が成功しました"})
}

// ゲストログイン
func GuestLogin(c *gin.Context) {
}

// ログイン
func Login(c *gin.Context) {
	var paramSignUp ParamSignUp

	// jsonからデータを受け取る
	if err := c.ShouldBindJSON(&paramSignUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "jsonからの変換に失敗しました"})
		return
	}

	// メールアドレスとパワスワードのバリデーションチェック
	if paramSignUp.invalid() == true {
		c.JSON(http.StatusBadRequest, gin.H{"message": "メールアドレスまたはパスワードが不適切です1"})
		return
	}

	// パラメータのチェック完了
	db := models.NewSQLHandler()

	// ログイン認証チェック
	var user User
	if err := db.DB.Table("mst_user").Where("msu_email = ?", paramSignUp.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// データが見つからなかった場合
			fmt.Println("user not found", paramSignUp.Email)
			c.JSON(http.StatusNotFound, gin.H{"error": err, "message": "アカウントが存在しないです"})
		} else {
			fmt.Println("database error", err.Error())
		}
		return
	}

	if err := crypto.CompareHashAndPassword(user.Password, paramSignUp.Password); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err, "message": "メールアドレスまたはパスワードが不適切です2"})
		return
	}

	// 自動ログインようにセッションを保存する

	c.JSON(http.StatusOK, gin.H{"message": "ログインが完了しました"})
}

func GetSession(c *gin.Context) {
	cookieKey := "session_id"
	redisValue := redis.GetSession(c, cookieKey)

	c.JSON(http.StatusOK, gin.H{"redisValue": redisValue})
}
