package service

import (
	"math/rand"
	"strconv"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/repository"
)

func FetchGachaResult(queryValue map[string][]string, err error) ([]model.Sushi, error) {
	var totalBudget int
	sushiList := fetchSushiData()
	if queryValue["value"] != nil {
		totalBudget, err = strconv.Atoi(queryValue["value"][0])
	}
	sushiPriceList := choiseSushiPriceCondition(sushiList, totalBudget)
	return sushiPriceList, err
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
