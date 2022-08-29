package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/service"
)

const DEFAULT_BUDGET int = 1000

func GachaPrice(w http.ResponseWriter, r *http.Request) {
	queryValue, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	sushiPriceList, err := service.FetchGachaResult(queryValue)
	if err != nil {
		log.Fatal(err)
	}
	sushiJson, err := json.Marshal(sushiPriceList)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sushiJson)
}
