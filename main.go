package main

import (
	"log"
	"net/http"

	controller "github.com/Sushiro-gacha/sushiro-gacha-api/app/controller"
)

// func HandllerHello() {
// 	return practicecontroller.RequestHello()
// }
func main() {
	StartMainServer()
}

func StartMainServer() {
	http.HandleFunc("/home/price/", controller.GathaPrice)
	//http.HandleFunc("/home/calorie/", controller.GachaCalories)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
