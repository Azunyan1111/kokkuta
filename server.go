package main

import (
	"github.com/Azunyan1111/kokkuta/contoroller"
	"net/http"
)

func main() {
	// リクエスト例
	// http://localhost:8080/api/send?body=hello_world2
	http.HandleFunc("/", contoroller.IndexHandler)        // メンイページ
	http.HandleFunc("/api/send", contoroller.SendHandler) // 送信ボタンのAPI
	http.HandleFunc("/api/good", contoroller.GoodHandler) // 送信ボタンのAPI

	http.HandleFunc("/request_token", contoroller.RequestTokenHandler) //apiを利用する時に使うリクエスト
	http.HandleFunc("/access_token", contoroller.AccessTokenHandler)   //apiを利用する時に使うアクストークン

	http.ListenAndServe(":8080", nil)
}
