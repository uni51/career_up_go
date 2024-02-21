package main

import (
	"fmt"
	"net/http"
	"trueserver/data"
	"trueserver/functions"
)

func add(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "*** はじめての関数 ***")

	// (2）自分で作ったモジュールを使った処理結果を出力
	result := functions.Add(5, 7)
	fmt.Fprintf(writer, "5+7=%d", result)
}

func sub(writer http.ResponseWriter, req *http.Request) {
	a, b := 5, 7

	output, result := functions.Sub(a, b)

	fmt.Fprintln(writer, "*** 複数の値を戻す関数 ***")
	fmt.Fprintf(writer, output, a, b, result)
}

func with_slices(writer http.ResponseWriter, req *http.Request) {
	sl_1 := []int{1, 2, 3, 4}

	fmt.Fprintln(writer, "*** スライスそのものを書き換える ***")

	// 処理前の値を確かめる
	fmt.Fprintln(writer, "\nsl_1は")
	fmt.Fprintln(writer, sl_1)

	functions.AddAll(sl_1, 9)

	fmt.Fprintln(writer, "\nいまやsl_1は")
	fmt.Fprintln(writer, sl_1)

	fmt.Fprintln(writer, "\n*** スライスをコピーして書き換える ***")

	sl_2 := functions.AddAndCopy(sl_1, 100)

	// 新しくできたsl_2
	fmt.Fprintln(writer, "\nsl_2は")
	fmt.Fprintln(writer, sl_2)

	// 元のsl_1
	fmt.Fprintln(writer, "\n一方sl_1は")
	fmt.Fprintln(writer, sl_1)
}

func with_structs(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "*** 構造体 ***")

	members := []data.Member{
		data.Member{"ゆみこ", 56, 1.24},
		data.Member{"トシオ", 44, 0.98},
		data.Member{"かおる", 70, 1.02},
	}

	fmt.Fprintln(writer, "*** 構造体を用いた関数 ***")
	fmt.Fprintln(writer, functions.DescribeAllMembers(members))

	fmt.Fprintln(writer, "*** 構造体を戻す関数 ***")
	fmt.Fprintln(writer, functions.DescribeMaxPointMember(members))

	fmt.Fprintln(writer, "*** 構造体のポインタを用いる関数 ***")

	member_add := &members[0]

	fmt.Fprintln(writer, functions.AddPointAndReport(&member_add, 12))
	fmt.Fprintln(writer, functions.Describe(members[0]))

	fmt.Fprintln(writer, "*** 構造体のコピーを返す関数 ***")
	friend, s_string := functions.CreateFriendAndReport(members[1], "エミコ")
	fmt.Fprintln(writer, s_string)

	fmt.Fprintln(writer, functions.Describe(friend))
	fmt.Fprintln(writer, functions.Describe(members[1]))

	members = append(members, friend)
	fmt.Fprintln(writer, "*** メソッドの使用 ***")
	fmt.Fprintln(writer, functions.DescribeM_AllMembers(members))

	fmt.Fprintln(writer, "\n<<お友達紹介特典>>")
	fmt.Fprintln(writer, functions.AddPointMAndReport(&members[1], 20))
	fmt.Fprintln(writer, functions.Describe(members[1]))
}

func with_pointers(writer http.ResponseWriter, req *http.Request) {
	mockmemory := []int{325, 14, 160, 440, 16, 175}

	fmt.Fprintln(writer, "\n<<アドレス[0]指定>>")
	fmt.Fprintln(writer, functions.DescribeMockStruct(mockmemory, 0))

	fmt.Fprintln(writer, "\n<<アドレス[3]指定>>")
	fmt.Fprintln(writer, functions.DescribeMockStruct(mockmemory, 3))

	fmt.Fprintln(writer, "\n*** ポインタを使う意味 ***")
	a, b := 10, 10

	aa := functions.UpdateOrCopy(a, &b)

	fmt.Fprintf(writer, "a=%d, b=%d, aa=%d", a, b, aa)
}

func with_methods(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "\n*** 連続処理 ***")

	marco := data.CreateTraveller("マルコ", 0, 0)
	marco = marco.Travel(2, 3).Travel(12, 24).Travel(45, 78).Goal()

	fmt.Fprintln(writer, marco.Record)

	fmt.Println(writer, "\n*** インターフェースの練習 ***")

	fractions := []data.Fraction{
		data.Half(1.5), data.Full(2), data.Half(2.5), data.Full(3), data.Half(3.5),
	}

	fmt.Fprintln(writer, functions.ShowFractions(fractions))

	fmt.Println(writer, "\n*** もっとそれらしいインターフェース ***")
	counters := []data.Counter{
		data.CharCounter{Content: "Let's count!"},
		data.CharCounter{Content: "一二三四五六七八九"},
		data.DigitCounter{Content: 2500},
		data.DigitCounter{Content: 1963061},
		data.CharCounter{Content: "以上！"},
	}

	fmt.Fprintln(writer, functions.CountAll(counters))
}

func main() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/with_slices", with_slices)
	http.HandleFunc("/with_structs", with_structs)
	http.HandleFunc("/with_pointers", with_pointers)
	http.HandleFunc("/with_methods", with_methods)

	http.ListenAndServe(":8090", nil)
}
