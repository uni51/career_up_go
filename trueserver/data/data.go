package data

import "fmt"

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
