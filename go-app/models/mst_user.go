package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLHandler ...
type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

// DBOpen は DB connectionを張る。
func DBOpen() {
	dbConn = NewSQLHandler()
}

// DBClose は DB connectionを張る。
func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {

	// TODO 環境変数に持つようにする必要がある
	user := "root"
	password := "root"
	host := "db"
	port := "3306"
	dbName := "book_recom"
	option := "charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Asia%2FTokyo"

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?" + option

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 接続確認
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	sqlDB, _ := db.DB()
	//コネクションプールの最大接続数を設定。
	sqlDB.SetMaxIdleConns(100)
	//接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	sqlDB.SetMaxOpenConns(100)
	//接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	sqlHandler := new(SQLHandler)
	db.Logger.LogMode(4)
	sqlHandler.DB = db

	return sqlHandler
}

// GetDBConn ...
func GetDBConn() *SQLHandler {
	return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

// Rollback ...
func RollBack() {
	dbConn.DB.Rollback()
}
