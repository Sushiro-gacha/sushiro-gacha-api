package main

import (
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/app/controller"
)

func main() {
	http.HandleFunc("/home/hello/", controller.RequestHello(w http.ResponseWriter, r *http.Request))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
