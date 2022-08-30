package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/service"
)

func GachaPrice(w http.ResponseWriter, r *http.Request) {
	queryValue, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	sushiPriceList, err := service.FetchGachaResult(queryValue)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	sushiJson, err := json.Marshal(sushiPriceList)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sushiJson)
}
