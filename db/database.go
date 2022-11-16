package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")
	dbconf := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4`, user, password, host, port, database_name)

	if Db, err = sql.Open("mysql", dbconf); err != nil {
		log.Fatal("Db open error:", err.Error())
	} else {
		fmt.Println("db Open")
	}
	Db.Begin()
	checkConnect(100)
	fmt.Println("db connected!!")

}

func checkConnect(count uint) {
	if err := Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		log.Printf("failed to ping by error '%#v'", &err)
		checkConnect(count)
	}
}
