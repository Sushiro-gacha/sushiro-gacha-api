package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/service"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

func GachaPrice(w http.ResponseWriter, r *http.Request) {
	sushiList := service.FetchSushiData()
	sushiJson, err := json.Marshal(sushiList)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(sushiJson)
}

// func GachaCalories(w http.ResponseWriter, r *http.Request) Calorie {

// }
