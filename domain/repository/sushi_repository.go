package repository

import (
	"database/sql"
	"log"

	"github.com/Sushiro-gacha/sushiro-gacha-api/domain/model"
)

// type sushiRepository struct {
// 	database.
// }

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

	// sushiEntityList := []model.SushiEntity{
	// 	{Id: 1, Category: "サーモンマウンテン", Name: "期間限定/エリア限定", Price: 330, Calorie: 211},
	// 	{Id: 2, Category: "えび天マウンテン", Name: "期間限定/エリア限定", Price: 330, Calorie: 383},
	// 	{Id: 3, Category: "天然インド鮪中落ちてんこ盛り", Name: "期間限定/エリア限定", Price: 165, Calorie: 73},
	// 	{Id: 4, Category: "うなきゅう巻", Name: "期間限定/エリア限定", Price: 165, Calorie: 199},
	// 	{Id: 5, Category: "大切り煮穴子", Name: "期間限定/エリア限定", Price: 165, Calorie: 66},
	// 	{Id: 6, Category: "大切りびんちょう", Name: "期間限定/エリア限定", Price: 110, Calorie: 97},
	// 	{Id: 7, Category: "大盛りねぎまぐろ", Name: "期間限定/エリア限定", Price: 110, Calorie: 126},
	// 	{Id: 8, Category: "大切りうなぎ", Name: "期間限定/エリア限定", Price: 110, Calorie: 73},
	// 	{Id: 9, Category: "うざく包み", Name: "期間限定/エリア限定", Price: 110, Calorie: 78},
	// 	{Id: 10, Category: "う巻きにぎり", Name: "期間限定/エリア限定", Price: 110, Calorie: 103},
	// 	{Id: 11, Category: "とろかつお", Name: "期間限定/エリア限定", Price: 110, Calorie: 99},
	// 	{Id: 12, Category: "大切り真いか", Name: "期間限定/エリア限定", Price: 110, Calorie: 76},
	// 	{Id: 13, Category: "天然インド鮪7貫食べ比べ", Name: "期間限定/エリア限定", Price: 1078, Calorie: 415},
	// 	{Id: 14, Category: "大切り甘鯛の天ぷら", Name: "期間限定/エリア限定", Price: 165, Calorie: 142},
	// 	{Id: 15, Category: "しまあじ", Name: "期間限定/エリア限定", Price: 165, Calorie: 51},
	// 	{Id: 16, Category: "いか梅しそにぎり", Name: "期間限定/エリア限定", Price: 165, Calorie: 72},
	// 	{Id: 17, Category: "まぐろビッグカツ", Name: "期間限定/エリア限定", Price: 110, Calorie: 151},
	// 	{Id: 18, Category: "トロあじ", Name: "期間限定/エリア限定", Price: 110, Calorie: 104},
	// 	{Id: 19, Category: "まぐろ・たまご", Name: "期間限定/エリア限定", Price: 110, Calorie: 98},
	// 	{Id: 20, Category: "サーモン・えび", Name: "期間限定/エリア限定", Price: 110, Calorie: 84},
	// 	{Id: 21, Category: "コーン・ツナ", Name: "期間限定/エリア限定", Price: 110, Calorie: 135},
	// 	{Id: 22, Category: "赤貝", Name: "期間限定/エリア限定", Price: 110, Calorie: 63},
	// 	{Id: 23, Category: "白とり貝バジルレモン", Name: "期間限定/エリア限定", Price: 110, Calorie: 83},
	// 	{Id: 24, Category: "ビーフ100％ハンバーグ", Name: "期間限定/エリア限定", Price: 110, Calorie: 135},
	// 	{Id: 25, Category: "ピリ辛イカキムチ軍艦", Name: "期間限定/エリア限定", Price: 110, Calorie: 78},
	// 	{Id: 26, Category: "梅きゅう巻", Name: "期間限定/エリア限定", Price: 110, Calorie: 151},
	// 	{Id: 27, Category: "季節のいなり（枝豆・ひじき）", Name: "期間限定/エリア限定", Price: 110, Calorie: 158},
	// 	{Id: 28, Category: "釜玉うどん", Name: "期間限定/エリア限定", Price: 330, Calorie: 283},
	// 	{Id: 29, Category: "もずくとあおさの赤だし", Name: "期間限定/エリア限定", Price: 220, Calorie: 47},
	// 	{Id: 30, Category: "もずくとあおさの味噌汁", Name: "期間限定/エリア限定", Price: 220, Calorie: 40},
	// 	{Id: 31, Category: "沖縄県産そでいかみみ唐揚げ", Name: "期間限定/エリア限定", Price: 330, Calorie: 234},
	// 	{Id: 32, Category: "フローズンフルーツインティー", Name: "期間限定/エリア限定", Price: 308, Calorie: 98},
	// 	{Id: 33, Category: "オレンジ＆ヨーグルトのアイスケーキ", Name: "期間限定/エリア限定", Price: 220, Calorie: 38},
	// 	{Id: 34, Category: "まぐろ", Name: "にぎり", Price: 110, Calorie: 78},
	// 	{Id: 35, Category: "漬けまぐろ", Name: "にぎり", Price: 110, Calorie: 81},
	// 	{Id: 36, Category: "サーモン", Name: "にぎり", Price: 110, Calorie: 95},
	// 	{Id: 37, Category: "オニオンサーモン", Name: "にぎり", Price: 110, Calorie: 124},
	// 	{Id: 38, Category: "焼とろサーモン", Name: "にぎり", Price: 110, Calorie: 95},
	// 	{Id: 39, Category: "おろし焼とろサーモン", Name: "にぎり", Price: 110, Calorie: 96},
	// 	{Id: 40, Category: "ジャンボとろサーモン", Name: "にぎり", Price: 110, Calorie: 67},
	// 	{Id: 41, Category: "サーモンちーず", Name: "にぎり", Price: 110, Calorie: 133},
	// 	{Id: 42, Category: "炙りサーモンバジルチーズ", Name: "にぎり", Price: 110, Calorie: 137},
	// 	{Id: 43, Category: "はまち", Name: "にぎり", Price: 110, Calorie: 111},
	// 	{Id: 44, Category: "焼き鯖", Name: "にぎり", Price: 110, Calorie: 189},
	// 	{Id: 45, Category: "〆真さば", Name: "にぎり", Price: 110, Calorie: 115},
	// 	{Id: 46, Category: "〆真さば(ごまネギ)", Name: "にぎり", Price: 110, Calorie: 117},
	// 	{Id: 47, Category: "〆いわし(ネギ・生姜)", Name: "にぎり", Price: 110, Calorie: 103},
	// 	{Id: 48, Category: "こはだ", Name: "にぎり", Price: 110, Calorie: 80},
	// 	{Id: 49, Category: "赤えび", Name: "にぎり", Price: 110, Calorie: 40},
	// 	{Id: 50, Category: "えび", Name: "にぎり", Price: 110, Calorie: 72},
	// 	{Id: 51, Category: "えびチーズ", Name: "にぎり", Price: 110, Calorie: 106},
	// 	{Id: 52, Category: "えびバジルチーズ", Name: "にぎり", Price: 110, Calorie: 110},
	// 	{Id: 53, Category: "えびアボカド", Name: "にぎり", Price: 110, Calorie: 116},
	// 	{Id: 54, Category: "甘えび", Name: "にぎり", Price: 110, Calorie: 65},
	// 	{Id: 55, Category: "煮あなご", Name: "にぎり", Price: 110, Calorie: 97},
	// 	{Id: 56, Category: "かにカマ天にぎり", Name: "にぎり", Price: 110, Calorie: 163},
	// 	{Id: 57, Category: "たまご", Name: "にぎり", Price: 110, Calorie: 117},
	// 	{Id: 58, Category: "豚塩カルビ", Name: "にぎり", Price: 110, Calorie: 116},
	// 	{Id: 59, Category: "生ハム", Name: "にぎり", Price: 110, Calorie: 112},
	// 	{Id: 60, Category: "黒門伊勢屋のわさびなす", Name: "にぎり", Price: 110, Calorie: 69},
	// 	{Id: 61, Category: "特ネタ中とろ", Name: "にぎり", Price: 165, Calorie: 91},
	// 	{Id: 62, Category: "びんとろ", Name: "にぎり", Price: 165, Calorie: 90},
	// 	{Id: 63, Category: "生サーモン", Name: "にぎり", Price: 165, Calorie: 72},
	// 	{Id: 64, Category: "サーモンバジルモッツァレラ", Name: "にぎり", Price: 165, Calorie: 132},
	// 	{Id: 65, Category: "たい", Name: "にぎり", Price: 165, Calorie: 93},
	// 	{Id: 66, Category: "漬けごま真鯛", Name: "にぎり", Price: 165, Calorie: 107},
	// 	{Id: 67, Category: "大えび", Name: "にぎり", Price: 165, Calorie: 45},
	// 	{Id: 68, Category: "ほたて貝柱", Name: "にぎり", Price: 165, Calorie: 76},
	// 	{Id: 69, Category: "大つぶ貝", Name: "にぎり", Price: 165, Calorie: 66},
	// 	{Id: 70, Category: "たこ", Name: "にぎり", Price: 165, Calorie: 72},
	// 	{Id: 71, Category: "うなぎの蒲焼き", Name: "にぎり", Price: 165, Calorie: 103},
	// 	{Id: 72, Category: "生ハムバジルモッツァレラ", Name: "にぎり", Price: 165, Calorie: 97},
	// 	{Id: 73, Category: "特ネタ大とろ", Name: "にぎり", Price: 330, Calorie: 79},
	// 	{Id: 74, Category: "特ネタ大とろ焦がし醤油", Name: "にぎり", Price: 330, Calorie: 80},
	// 	{Id: 75, Category: "大切りあわび2貫", Name: "にぎり", Price: 330, Calorie: 70},
	// 	{Id: 76, Category: "軍艦ねぎまぐろ", Name: "軍艦・巻物", Price: 110, Calorie: 119},
	// 	{Id: 77, Category: "まぐろ山かけ", Name: "軍艦・巻物", Price: 110, Calorie: 94},
	// 	{Id: 78, Category: "まぐろユッケ(卵黄醤油)", Name: "軍艦・巻物", Price: 110, Calorie: 115},
	// 	{Id: 79, Category: "まぐたく", Name: "軍艦・巻物", Price: 110, Calorie: 112},
	// 	{Id: 80, Category: "かにみそ", Name: "軍艦・巻物", Price: 110, Calorie: 94},
	// 	{Id: 81, Category: "軍艦甘えび", Name: "軍艦・巻物", Price: 110, Calorie: 100},
	// 	{Id: 82, Category: "とびこ軍艦", Name: "軍艦・巻物", Price: 110, Calorie: 79},
	// 	{Id: 83, Category: "たらマヨ", Name: "軍艦・巻物", Price: 110, Calorie: 142},
	// 	{Id: 84, Category: "たらこ", Name: "軍艦・巻物", Price: 110, Calorie: 88},
	// 	{Id: 85, Category: "コーン", Name: "軍艦・巻物", Price: 110, Calorie: 119},
	// 	{Id: 86, Category: "ツナサラダ", Name: "軍艦・巻物", Price: 110, Calorie: 150},
	// 	{Id: 87, Category: "シーサラダ", Name: "軍艦・巻物", Price: 110, Calorie: 119},
	// 	{Id: 88, Category: "カニ風サラダ", Name: "軍艦・巻物", Price: 110, Calorie: 127},
	// 	{Id: 89, Category: "いかオクラめかぶ", Name: "軍艦・巻物", Price: 110, Calorie: 72},
	// 	{Id: 90, Category: "めかぶ長芋納豆軍艦", Name: "軍艦・巻物", Price: 110, Calorie: 91},
	// 	{Id: 91, Category: "小粒納豆", Name: "軍艦・巻物", Price: 110, Calorie: 97},
	// 	{Id: 92, Category: "きゅうり巻", Name: "軍艦・巻物", Price: 110, Calorie: 139},
	// 	{Id: 93, Category: "新香巻", Name: "軍艦・巻物", Price: 110, Calorie: 160},
	// 	{Id: 94, Category: "海老フライアボカドロール", Name: "軍艦・巻物", Price: 110, Calorie: 131},
	// 	{Id: 95, Category: "いくら", Name: "軍艦・巻物", Price: 165, Calorie: 82},
	// 	{Id: 96, Category: "軍艦馬刺しねぎとろ", Name: "軍艦・巻物", Price: 165, Calorie: 95},
	// 	{Id: 97, Category: "鉄火巻", Name: "軍艦・巻物", Price: 165, Calorie: 169},
	// 	{Id: 98, Category: "コク旨まぐろ醤油ラーメン", Name: "サイドメニュー", Price: 385, Calorie: 328},
	// 	{Id: 99, Category: "鯛だし塩ラーメン", Name: "サイドメニュー", Price: 385, Calorie: 236},
	// 	{Id: 100, Category: "濃厚えび味噌ワンタンメン", Name: "サイドメニュー", Price: 385, Calorie: 346},
	// 	{Id: 101, Category: "わかめうどん", Name: "サイドメニュー", Price: 330, Calorie: 199},
	// 	{Id: 102, Category: "かけうどん", Name: "サイドメニュー", Price: 165, Calorie: 177},
	// 	{Id: 103, Category: "魚のアラの赤だし", Name: "サイドメニュー", Price: 198, Calorie: 113},
	// 	{Id: 104, Category: "あさりの赤だし", Name: "サイドメニュー", Price: 198, Calorie: 73},
	// 	{Id: 105, Category: "あさりの味噌汁", Name: "サイドメニュー", Price: 198, Calorie: 66},
	// 	{Id: 106, Category: "茶碗蒸し", Name: "サイドメニュー", Price: 220, Calorie: 76},
	// 	{Id: 107, Category: "たこの唐揚げ", Name: "サイドメニュー", Price: 308, Calorie: 249},
	// 	{Id: 108, Category: "なんこつ唐揚げ", Name: "サイドメニュー", Price: 165, Calorie: 239},
	// 	{Id: 109, Category: "フライドポテト", Name: "サイドメニュー", Price: 132, Calorie: 189},
	// 	{Id: 110, Category: "かぼちゃの天ぷら", Name: "サイドメニュー", Price: 132, Calorie: 74},
	// 	{Id: 111, Category: "店内仕込の海鮮ポテサラ(ガリ入)", Name: "サイドメニュー", Price: 132, Calorie: 169},
	// 	{Id: 112, Category: "生ビール　ジョッキ", Name: "ドリンク", Price: 528, Calorie: 235},
	// 	{Id: 113, Category: "生ビール　グラス", Name: "ドリンク", Price: 385, Calorie: 235},
	// 	{Id: 114, Category: "大吟醸", Name: "ドリンク", Price: 638, Calorie: 540},
	// 	{Id: 115, Category: "生貯蔵酒", Name: "ドリンク", Price: 418, Calorie: 440},
	// 	{Id: 116, Category: "角ハイボール", Name: "ドリンク", Price: 385, Calorie: 275},
	// 	{Id: 117, Category: "こだわり酒場のレモンサワー", Name: "ドリンク", Price: 385, Calorie: 240},
	// 	{Id: 118, Category: "オールフリー（ﾉﾝｱﾙｺｰﾙﾋﾞｰﾙ）", Name: "ドリンク", Price: 385, Calorie: 0},
	// 	{Id: 119, Category: "りんごジュース　国産100％果汁", Name: "ドリンク", Price: 165, Calorie: 80},
	// 	{Id: 120, Category: "アイスコーヒー", Name: "ドリンク", Price: 165, Calorie: 14},
	// 	{Id: 121, Category: "ホットコーヒー", Name: "ドリンク", Price: 165, Calorie: 24},
	// 	{Id: 122, Category: "アイスカフェラテ", Name: "ドリンク", Price: 198, Calorie: 224},
	// 	{Id: 123, Category: "ホットカフェラテ", Name: "ドリンク", Price: 198, Calorie: 283},
	// }
	return sushiEntityList
}
