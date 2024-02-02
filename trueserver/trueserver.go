package main

import (
	"fmt"
	"net/http"
	"trueserver/functions"
)

func add(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "*** はじめての関数 ***")

	// (2）自分で作ったモジュールを使った処理結果を出力
	result := functions.Add(5, 7)
	fmt.Fprintf(writer, "5+7=%d", result)
}

func main() {
	http.HandleFunc("/add", add)

	http.ListenAndServe(":8090", nil)
}
