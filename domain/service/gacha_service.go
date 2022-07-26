package service

import (
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
	repository "github.com/Sushiro-gacha/sushiro-gacha-api/domain/repository"
)

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
