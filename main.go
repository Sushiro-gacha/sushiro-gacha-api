package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/app/controller"
	database "github.com/Sushiro-gacha/sushiro-gacha-api/db"
)

func main() {

	db := database.Connect()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	err := db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("データベース接続成功")
	}
	// StartMainServer()
}

func StartMainServer() {
	http.HandleFunc("/gacha/price", controller.GachaPrice)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
