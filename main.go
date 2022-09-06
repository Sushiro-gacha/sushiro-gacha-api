package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/app/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	_, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	// StartMainServer()
}
func StartMainServer() {
	http.HandleFunc("/gacha/price", controller.GachaPrice)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "test_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
