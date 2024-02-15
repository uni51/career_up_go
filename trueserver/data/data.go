package data

import "fmt"

type Member struct {
	Name  string
	Point int
	Coeff float64 // 係数
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

func CreateFriendMember(member Member, name string) Member {
	member.Name = name
	return member
}
