# sushiro-gacha-api

## ディレクトリ構造
```
sushiro-gacha-api
├── app
│   └── controller
│       └── gachaController.go  // ガチャ用のコントローラー
├── db
│   └── db.go                   // DB接続用のクラス
├── domain
│   ├── model
│   │   └── sushi.go            // 寿司の情報を登録する構造体を定義
│   ├── repository
│   │   └── sushiRepository.go  // 寿司テーブル接続用のリポジトリ
│   └── service
│       └── gachaService.go     // ガチャ用のサービスクラス
├── go.mod
└── main.go                     // メインクラス 今回はルーティングも兼ねる
```
