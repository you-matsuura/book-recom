package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func TestConnect() {
	fmt.Println("パッケージの読み込み成功")
}

func ConnectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error loading location:", err)
		return nil
	}
	c := mysql.Config{
		DBName:    "book_recom",
		User:      "root",
		Passwd:    "root",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error opening database:", err)
		return nil
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	// GetData(db)

	// ここで、正常に接続されたデータベースを返します。
	return db
}


func GetData(db *sql.DB) {
	// 
	type User struct {
		Id int
		UserName string
	}

	defer db.Close()

	// _, err = db.Exec(`SELECT * form mst_user`)

	rows, err := db.Query("SELECT msu_user_id, msu_nick_name FROM mst_user")

	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}

	defer rows.Close()

	for rows.Next() {

		// var user User

		user := &User{}

		err := rows.Scan(&user.Id, &user.UserName)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Id, user.UserName)
	}
} 