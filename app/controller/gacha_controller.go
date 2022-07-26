package controller

// リクエストとレスポンスを司る

import (
	"encoding/json"
	"log"
	"net/http"

	service "github.com/Sushiro-gacha/sushiro-gacha-api/domain/service"
)

// func formatter_json(data []byte) []byte {
// 	var sushi []model.Sushi
// 	if err := json.Unmarshal(data, &sushi); err != nil {
// 		log.Fatal(err)
// 	}
// 	json_data, err := json.Marshal(sushi)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return json_data
// }

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
