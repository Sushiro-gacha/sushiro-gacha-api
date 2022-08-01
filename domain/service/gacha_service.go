package service

import (
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/repository"
)

func ChoiseSushiPriceCondition(sushiList []model.Sushi, maxPrice int) []model.Sushi {
	var resultSushiList []model.Sushi
	for sumPrice := 0; sumPrice < maxPrice; {
		resultSushiList = append(resultSushiList, sushiList[randomNumber])
		sumPrice = sumPrice + sushiList[randomNumber].Price
	}
	return resultSushiList
}

func FetchSushiData() []model.Sushi {
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
