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

func Describe(m Member) string {
	return fmt.Sprintf("%s さんのポイントは %d 点、有効ポイントは %.2f 点です", m.Name, m.Point, Effective(m))
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
