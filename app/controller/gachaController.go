package controller

// リクエストとレスポンスを司る

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Sushi struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func formatter_json(data []byte) []byte {
	var sushi []Sushi
	if err := json.Unmarshal(data, &sushi); err != nil {
		log.Fatal(err)
	}
	json_data, err := json.Marshal(sushi)
	if err != nil {
		log.Fatal(err)
	}
	return json_data
}

func get_sushi() []byte {
	jsonFromFile, err := ioutil.ReadFile("./app/controller/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	res := formatter_json(jsonFromFile)

	return res
}

func GathaPrice(w http.ResponseWriter, r *http.Request) {
	sushi_json := get_sushi()

	w.Header().Set("Content-Type", "application/json")
	w.Write(sushi_json)
}

// func GachaCalories(w http.ResponseWriter, r *http.Request) Calorie {

// }
