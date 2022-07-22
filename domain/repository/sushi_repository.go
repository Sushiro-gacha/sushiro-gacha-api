package repository

import (
	"encoding/json"
	"log"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

func return_sushi() model.SushiEntity {
	sushi_string := `{
			"id":       1,
			"category": "nigiri",
			"name":     "maguro",
			"proce":    100,
			"calories": 100,
		}`
	var sushi model.SushiEntity
	if err := json.Unmarshal([]byte(sushi_string), &sushi); err != nil {
		log.Fatal(err)
	}

	return sushi
}
