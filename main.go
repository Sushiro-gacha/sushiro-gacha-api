package main

import (
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/app/controller"
)

func main() {
	StartMainServer()
}

func StartMainServer() {
	http.HandleFunc("/gacha/price", controller.GachaPrice)
	http.HandleFunc("/testdb", controller.TestDbConnect)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
