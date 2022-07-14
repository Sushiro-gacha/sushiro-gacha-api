package controller

// リクエストとレスポンスを司る

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

func formatter_json(data []byte) []byte {
	var sushi []model.Sushi
	if err := json.Unmarshal(data, &sushi); err != nil {
		log.Fatal(err)
	}
	json_data, err := json.Marshal(sushi)
	if err != nil {
		log.Fatal(err)
	}
	return json_data
}

func GachaPrice(w http.ResponseWriter, r *http.Request) {
	jsonFromFile, err := ioutil.ReadFile("./app/controller/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	sushi_json := formatter_json(jsonFromFile)

	w.Header().Set("Content-Type", "application/json")
	w.Write(sushi_json)
}

// func GachaCalories(w http.ResponseWriter, r *http.Request) Calorie {

// }
