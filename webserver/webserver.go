package main

import (
	"fmt"
	"math"
	"net/http"
)

func hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "レッツゴー")
}

func algebra(writer http.ResponseWriter, req *http.Request) {
	a, b := 7, 8
	result := "%d + %d = %d\n"             // Fprintfの最初の引数に渡す
	fmt.Fprintf(writer, "***整数の足し算***\n")  // これはただの見出し
	fmt.Fprintf(writer, result, a, b, a+b) // 整数を文字列に埋め込んで出力

	c := 15.0
	d := float64(a)                 // 小数に変換
	result = "%.1f / %.1f = %.3f\n" // 表示桁数を指定する小数の書式
	fmt.Fprintf(writer, "\n***小数の割り算***\n")
	fmt.Fprintf(writer, result, c, d, c/d)
}

func mathfuncs(writer http.ResponseWriter, req *http.Request) {
	pow28 := int(math.Pow(2, 8)) // mathパッケージのPow関数（2の8乗）
	fmt.Fprintf(writer, "%dの%d乗は%d\n", 2, 8, pow28)

	rad30 := 30.0 * math.Pi / 180.0 //度をラジアンに変換
	fmt.Fprintf(writer, "\nsin%d°は %.3f\n", 30, math.Sin(rad30))
}

func arrays(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "***要素が5個の配列を定義***")
	arr1 := [5]int{2, 4, 6, 8, 10}
	fmt.Fprintln(writer, arr1)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/algebra", algebra)
	http.HandleFunc("/math", mathfuncs)
	http.HandleFunc("/arrays", arrays)
	http.ListenAndServe(":8090", nil)
}
