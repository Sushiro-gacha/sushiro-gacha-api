package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// url
const SSR_MENU_TOP = "https://www.akindo-sushiro.co.jp/menu/"

// section number
const MIN_SECTION = 1
const MAX_SECTION = 7

// regrexp
const REGEXP_PATTERN = "[0-9]+"

// pattern
const ALCHOL_PATTERN = "100gあたり"
const COFFEE_PATTERN = "砂糖なしの場合"

//calorie rate
// 500mlとして扱う
const ALCHOL_RATE float64 = 5.0

// 350mlとして扱う
const COFFEE_RATE float64 = 3.5

// デフォルト
const OTHER_RATE float64 = 1.0

type Sushi struct {
	Name     string
	Category string
	Price    int
	Calorie  int
}

func main() {
	fetchSushiro()
}

func fetchSushiro() {

	log.Println("start to fetch")
	doc, err := goquery.NewDocument(SSR_MENU_TOP)
	if err != nil {
		log.Fatalf("error!!")
	}

	// log.Println(doc.Html())
	sushiList := []Sushi{}
	for i := MIN_SECTION; i < MAX_SECTION; i++ {
		// 取得コンテンツを指定(02は使用されていない)
		var sectionName = "#anchor-sec0" + strconv.Itoa(i)

		doc.Find(sectionName).Each(func(_ int, categories *goquery.Selection) {
			// カテゴリ名を取得
			// 文頭に改行コードがあるのでトリムする(参考:https://devlights.hatenablog.com/entry/2020/10/29/174701)
			categoryName := strings.TrimLeft(categories.Find(".sec-ttl > a.acc-trigger").Text(), "\n")
			categories.Find("ul.item-list > li").Each(func(_ int, items *goquery.Selection) {
				// 食べ物の名前を取得
				sushiName := items.Find(".ttl").Text()
				//
				sushiPriceAndCal := items.Find(".price").Text()
				price, calorie, err := generatePriceAndCal(sushiPriceAndCal)
				if err != nil {
					log.Fatalf("error!!", err)
					return
				}

				sushi := Sushi{
					Name:     sushiName,
					Category: categoryName,
					Price:    price,
					Calorie:  calorie,
				}
				log.Printf("SUSHI data.	Name :  %s, Category :  %s, Price :  %d, Calorie :  %d", sushi.Name, sushi.Category, sushi.Price, sushi.Calorie)
				sushiList = append(sushiList, sushi)
			})
		})
	}
	log.Println(len(sushiList))
	log.Println("end to fetch")

}

func generatePriceAndCal(val string) (price int, calorie int, err error) {
	calorieRatio, formatedStr := calcRateAndFormatString(val)

	rex := regexp.MustCompile(REGEXP_PATTERN)
	strVal := rex.FindAllString(formatedStr, -1)
	if len(strVal) != 3 {
		return -1, -1, errors.New("期待したフォーマットになっていません")
	}

	priceStr := strVal[1]
	price, err = strconv.Atoi(priceStr)
	if err != nil {
		log.Fatalf("error!!")
		return -1, -1, err
	}
	calorieStr := strVal[2]
	calorieNoRate, err := strconv.Atoi(calorieStr)
	if err != nil {
		log.Fatalf("error!!")
		return -1, -1, err
	}
	calorie = int(float64(calorieNoRate) * calorieRatio)

	return price, calorie, nil

}

func calcRateAndFormatString(val string) (calorieRatio float64, formatedStr string) {
	// パターンによってカロリーの計算レートを変える
	// また、「100あたり~」のような数字を含む表現もあるのでreplaceで消しておく
	var replacedStr string
	if strings.Contains(val, ALCHOL_PATTERN) {
		calorieRatio = ALCHOL_RATE
		replacedStr = strings.ReplaceAll(val, ALCHOL_PATTERN, "")
	} else if strings.Contains(val, COFFEE_PATTERN) {
		calorieRatio = COFFEE_RATE
		replacedStr = strings.ReplaceAll(val, COFFEE_PATTERN, "")
	} else {
		calorieRatio = OTHER_RATE
		replacedStr = val
	}

	// 4桁以上の金額のカンマつき金額対策でカンマを消しておく
	formatedStr = strings.ReplaceAll(replacedStr, ",", "")

	return calorieRatio, formatedStr
}
