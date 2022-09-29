package service

import (
	"math/rand"
	"strconv"

	database "github.com/Sushiro-gacha/sushiro-gacha-api/db"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/repository"
)

func FetchGachaResult(queryValue map[string][]string) ([]model.Sushi, error) {
	var err error
	var totalBudget int
	sushiList := fetchSushiData()
	if queryValue["value"] != nil {
		totalBudget, err = strconv.Atoi(queryValue["value"][0])
	}
	sushiPriceList := choiseSushiPriceCondition(sushiList, totalBudget)
	return sushiPriceList, err
}

func choiseSushiPriceCondition(sushiList []model.Sushi, totalBudget int) []model.Sushi {
	resultSushiList := []model.Sushi{}
	restBudget := totalBudget
	//Put sushiList's childs which is within rest of total budget into underValueSushiList until there are no more childs to add.
	for {
		underValueSushiList := []model.Sushi{}
		for _, sushi := range sushiList {
			if sushi.Price <= restBudget {
				underValueSushiList = append(underValueSushiList, sushi)
			}
		}
		if len(underValueSushiList) == 0 {
			break
		}
		//Choise one child from underValueSushiList and put it into resultList
		randomNumber := rand.Intn(len(underValueSushiList))
		sushiChild := underValueSushiList[randomNumber]
		resultSushiList = append(resultSushiList, sushiChild)
		//Substract Chiled's cost from rest-budget
		restBudget = restBudget - sushiChild.Price
	}
	return resultSushiList
}

func fetchSushiData() []model.Sushi {
	sushiList := []model.Sushi{}
	sushiEntityList := repository.FetchSushiData(database.Db)
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

func TestFetchSushi() []model.SushiEntity {
	testSushiEntity := repository.FetchSushiData(database.Db)
	return testSushiEntity
}
