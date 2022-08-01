package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/service"
)

func GachaPrice(w http.ResponseWriter, r *http.Request) {
	sushiList := service.FetchSushiData()
	queryMap := r.URL.Query()
	value, _ := strconv.Atoi(queryMap["value"][0])
	sushiPriceList := service.ChoiseSushiPriceCondition(sushiList, value)
	sushiJson, err := json.Marshal(sushiPriceList)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(sushiJson)
}

// func GachaCalories(w http.ResponseWriter, r *http.Request) Calorie {

// }
