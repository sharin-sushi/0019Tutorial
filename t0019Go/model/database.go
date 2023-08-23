package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// var Dbsql *sql.DB //テスト用

func init() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")

	var path string = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)

	// sqlへ接続するための文字列の生成
	fmt.Println("path=" + path + "in database.go")
	var err error

	// fmt.Printf("%s\n%s\n", path, err)

	//素の書き方
	// if Db, err = sql.Open("mysql", path); err != nil {
	// 	fmt.Printf("database.goのinitでエラー発生:err=%s, path=%s", err, path)
	// 	// log.Fatal("Db open error:", err.Error())
	// }

	// Db, err := gorm.Open(sqlite.Open(path), &gorm.Config{}) //記事通りだとMySQLへの接続でも"sqlite""
	Db, err = gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	if Db == nil {
		panic("failed to connect database")

	} //このif Db文消したい意味的に重複してる

	fmt.Printf("%s\n%s\n", path, err)
}

// user := "ユーザー名"
// password := "パスワード"
// host := "ホスト名"
// port := "ポート番号"
// database := "データベース名"
