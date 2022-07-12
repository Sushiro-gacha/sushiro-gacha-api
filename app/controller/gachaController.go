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

type Calorie struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Calorie int    `json:"calorie"`
}

func get_sushi() []byte {
	jsonFromFile, err := ioutil.ReadFile("app/controller/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(jsonFromFile)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func get_calorie() Calorie {
	jsonFromFile, err := ioutil.ReadFile("./sample2.json")
	if err != nil {
		log.Fatal(err)
	}

	var jsonData Calorie
	err = json.Unmarshal(jsonFromFile, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

func GathaPrice(w http.ResponseWriter, r *http.Request) {
	sushi_json := get_sushi()

	w.Header().Set("Content-Type", "application/json")
	w.Write(sushi_json)
}

func GachaCalories(w http.ResponseWriter, r *http.Request) Calorie {
	calorie_json := get_calorie()

	return calorie_json
}
