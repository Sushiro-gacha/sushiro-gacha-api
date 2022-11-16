package repository

import (
	"database/sql"
	"log"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

func FetchSushiData(db *sql.DB) []model.SushiEntity {
	sushiEntityList := []model.SushiEntity{}
	rows, err := db.Query("SELECT * FROM sushi")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}

	for rows.Next() {
		s := model.SushiEntity{}
		if err := rows.Scan(&s.Id, &s.Category, &s.Name, &s.Price, &s.Calorie); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		sushiEntityList = append(sushiEntityList, s)
	}

	return sushiEntityList
}
