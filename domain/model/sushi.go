package model

// ここにmodelの情報を定義する

type Sushi struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Calorie  int    `json:"calorie"`
}
