package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

// ここにDBの接続メソッドを定義する。
// イメージとしては、起動時に「func init()」で接続を定義して
// repositoryからここの接続情報を呼び出す

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
