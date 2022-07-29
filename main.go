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
	http.HandleFunc("/gacha/price/", controller.GachaPrice)
	//http.HandleFunc("/home/calorie/", controller.GachaCalories)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
