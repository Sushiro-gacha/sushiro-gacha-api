package service

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/repository"
)

func FetchGachaResult(w http.ResponseWriter, r *http.Request) []model.Sushi {
	var totalBudget int
	sushiList := fetchSushiData()
	queryMap, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	if queryMap["value"] != nil {
		totalBudget, err = strconv.Atoi(queryMap["value"][0])
	}
	if err != nil {
		log.Fatal(err)
	}
	//ここまで通ってる
	sushiPriceList := choiseSushiPriceCondition(sushiList, totalBudget)
	// output: [0/0]0x0
	println(sushiPriceList)
	return sushiPriceList
}

func choiseSushiPriceCondition(sushiList []model.Sushi, totalBudget int) []model.Sushi {
	var resultSushiList []model.Sushi
	for sumPrice := 0; sumPrice > totalBudget; {
		randomNumber := rand.Intn(len(sushiList)) + 1
		sushiChild := sushiList[randomNumber]
		sumPrice = sumPrice + sushiChild.Price
		if sumPrice < totalBudget {
			resultSushiList = append(resultSushiList, sushiChild)
		}
	}
	return resultSushiList
}

func fetchSushiData() []model.Sushi {
	var sushiList []model.Sushi
	sushiEntityList := repository.FetchSushiData()
	for _, entity := range sushiEntityList {
		sushi := model.Sushi{}
		sushi.Id = entity.Id
		sushi.Category = entity.Category
		sushi.Name = entity.Name
		sushi.Price = entity.Price
		sushi.Calorie = entity.Calorie

		sushiList = append(sushiList, sushi)
	}
	return sushiList
}
