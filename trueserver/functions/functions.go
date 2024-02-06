package functions

import (
	"fmt"
	"trueserver/data"
)

func Add(a int, b int) int {
	return a + b
}

func Sub(a, b int) (string, int) {
	return "%d-%d は %d なのだ", a - b
}

func AddAll(sl []int, a int) {
	for i := 0; i < len(sl); i++ {
		sl[i] += a
	}
}

func AddAndCopy(sl []int, a int) []int {
	sl_cp := []int{}
	for i := 0; i < len(sl); i++ {
		sl_cp = append(sl_cp, sl[i]+a)
	}
	return sl_cp
}

func Describe(member data.Member) string {
	s_string := fmt.Sprintf(data.Describe(member))

	s_string += "\n"
	return s_string
}

func DescribeAllMembers(members []data.Member) string {
	s_string := ""
	for _, v := range members {
		s_string += Describe(v)
	}
	return s_string
}

func DescribeMaxPointMember(members []data.Member) string {
	s_string := "有効ポイント最大の方は\n"

	mpm := data.MaxPointMember(members)

	s_string += fmt.Sprintf("%s さん\n", mpm.Name)
	s_string += "\n"

	return s_string
}
