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

	fmt.Fprintln(writer, "\n***要素5個で定義した配列に3つの要素しか定義しないとどうなる？***")
	arr2 := [5]int{1, 3, 5}
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n***要素の値は変更可？***")
	arr2[4] = 99
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n***配列の一部を参照するスライス***")
	sl1 := arr1[1:3]
	sl2 := arr2[3:]
	fmt.Fprintln(writer, sl1)
	fmt.Fprintln(writer, sl2)

	fmt.Fprintln(writer, "\n***スライスの値を変更するとどうなる？***")
	sl1[1] = 36
	fmt.Fprintln(writer, sl1)
	fmt.Fprintln(writer, arr1) // 元の配列の値も変わる
}

func slices(writer http.ResponseWriter, req *http.Request) {
	sl := []int{30, 45, 60, 90, 180}

	var rad_v float64 //初期値を渡さず、型だけで変数を定義

	for _, v := range sl { //インデックスと値を取得
		rad_v = float64(v) * math.Pi / 180.0 //変数rad_vの値を置き換えていく
		fmt.Fprintf(writer, "sin%d°は %.3f\n\n",
			v, math.Sin(rad_v)) //そのたびに出力する

	}

	fmt.Fprintln(writer, "\n***スライスなら要素を増やせる***")
	sl = append(sl, 225, 275, 360)
	fmt.Fprintln(writer, sl)

	fmt.Fprintln(writer, "\n***スライスの一部を参照するスライス***")
	sl_sl := sl[2:5]
	fmt.Fprintln(writer, sl_sl)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/algebra", algebra)
	http.HandleFunc("/math", mathfuncs)
	http.HandleFunc("/arrays", arrays)
	http.HandleFunc("/slices", slices)

	http.ListenAndServe(":8090", nil)
}
