package controller

import "net/http"

// For Practice.

func get_hello() {

}

//THis function get "Hello World" json object.
func RequestHello(w http.ResponseWriter, r *http.Request) {
	hello_json := get_hello()
	return hello_json
}

// リクエストとレスポンスを司る

// func GathaPrice() {
// 	// 値段でガチャを引く場合
// }

// func GachaCalories() {
// 	// カロリーででガチャを引く場合
// }
