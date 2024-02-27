package data

import (
	"fmt"
	"math"
)

type Member struct {
	Name  string
	Point int
	Coeff float64 // 係数
}

type Traveller struct {
	Name   string
	X      int
	Y      int
	Record string
}

type Half float64

type Full int

type Fraction interface {
	Value() string
}

type Counter interface {
	DoCount() string
}

type CharCounter struct {
	Content string
}

type DigitCounter struct {
	Content int
}

type MockReader interface {
	Read(content string)
	Write() string
}

type StringReader struct {
	Memory string
}

type IntReader struct {
	Memory []int
}

func (reader *StringReader) Read(content string) {
	reader.Memory += content
	reader.Memory += "\n"
}

func (reader *IntReader) Read(content string) {
	digits := "0123456789"
	for _, v := range content {
		for i, s := range digits {
			if v == s {
				reader.Memory = append(reader.Memory, i)
				break
			}
		}
	}
}

func (reader StringReader) Write() string {
	s_string := "StringReaderインスタンスの中身は\n"
	s_string += "「"
	s_string += reader.Memory
	s_string += "」"
	return s_string
}

func (reader IntReader) Write() string {
	s_string := "IntReaderインスタンスの中身は\n"
	s_string += "["
	for _, v := range reader.Memory {
		s_string += fmt.Sprintf("%d ", v)
	}
	s_string += "]"
	return s_string
}

func (reader IntReader) Render2Int() int {
	sum := 0
	memory := reader.Memory
	lm := len(memory)
	for i := 0; i < lm; i++ {
		mag := math.Pow10(lm - i)
		sum += memory[i] * int(mag)
	}
	return sum / 10
}

func (counter CharCounter) DoCount() string {
	content := counter.Content                // (1) 処理する値を取り出す
	s_string := fmt.Sprintf("「%s」は", content) // (2) まず、内容を表示する

	// (3) 文字列をUnicode記号のスライスに変換して、その長さを取得
	s_string += fmt.Sprintf(" %d 文字です", len([]rune(content)))

	return s_string
}

func (counter DigitCounter) DoCount() string {
	content := counter.Content
	content_str := fmt.Sprintf("%d", content) // (1) 整数を文字列に変換
	s_string := fmt.Sprintf("「%d」は", content) // (2) 何の桁数か明示

	// (3) 文字列をUnicode記号のスライスに変換して、その長さを取得
	s_string += fmt.Sprintf(" %d 桁です", len([]rune(content_str)))

	return s_string
}

func (h Half) Value() string {
	return fmt.Sprintf("%.1f", float64(h))
}

func (f Full) Value() string {
	return fmt.Sprintf("%d", int(f))
}

func CreateTraveller(name string, x int, y int) Traveller {
	t := Traveller{}
	t.Name = name
	t.X = x
	t.Y = y
	t.Record = fmt.Sprintf("%s は (%d, %d) 地点よりスタート\n", t.Name, t.X, t.Y)
	return t
}

func (t Traveller) Travel(x int, y int) Traveller {
	t.X = x
	t.Y = y
	t.Record += fmt.Sprintf("(%d, %d) へ移動\n", t.X, t.Y)
	return t
}

func (t Traveller) Goal() Traveller {
	t.Record += "到着です\n"
	return t
}

func Effective(m Member) float64 {
	return float64(m.Point) * m.Coeff
}

func (member Member) EffectiveM() float64 {
	return float64(member.Point) * member.Coeff
}

func Describe(member Member) string {
	return fmt.Sprintf("%s さんのポイントは %d 点、有効ポイントは %.2f 点です", member.Name, member.Point, Effective(member))
}

func (member Member) DescribeM() string {
	return fmt.Sprintf("%s さんのポイントは %d 点、有効ポイントは %.2f 点です", member.Name, member.Point, member.EffectiveM())
}

// 有効ポイントが最大のメンバーを返す
func MaxPointMember(members []Member) Member {

	mpm := members[0]

	for _, v := range members {
		if Effective(v) > Effective(mpm) {
			mpm = v
		}
	}

	return mpm
}

func AddPoint(member **Member, p int) {
	(**member).Point += p
}

func (member *Member) AddPointM(p int) {
	member.Point += p
}

func CreateFriendMember(member Member, name string) Member {
	member.Name = name
	return member
}
