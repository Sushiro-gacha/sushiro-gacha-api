package main

func main() {
	StartMainServer()
}

func StartMainServer() {
	// ここでルーティングを行う
	// controllerのメソッドをhttp.Hundlefuncで登録してみよう
	// 公式：https://pkg.go.dev/net/http

	// errorを返すよう設定
	// return http.ListenAndServe(":8080", nil)
}
