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
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
