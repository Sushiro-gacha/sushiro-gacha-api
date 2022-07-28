package repository

import (
	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

func FetchSushiData() []model.SushiEntity {
	sushi := []model.SushiEntity{{Id: 1, Category: "nigiri", Name: "maguro", Price: 100, Calorie: 100}, {Id: 2, Category: "nigiri", Name: "samon", Price: 120, Calorie: 200}}

	return sushi
}
