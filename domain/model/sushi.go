package model

type Sushi struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Calorie  int    `json:"calorie"`
}
