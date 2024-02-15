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

func DescribeM_AllMembers(members []data.Member) string {
	s_string := "メソッドを使って書き出しても\n"
	for _, v := range members {
		s_string += v.DescribeM()
		s_string += "\n"
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

func DescribeMockStruct(mockmemory []int, mockaddress int) string {
	s_string := fmt.Sprintf("名前は %d さん、", mockmemory[mockaddress])
	s_string += fmt.Sprintf("年齢は %d 歳、", mockmemory[mockaddress+1])
	s_string += fmt.Sprintf("身長は %d cm", mockmemory[mockaddress+2])
	return s_string
}

func UpdateOrCopy(a int, b *int) int {
	a += 3
	*b += 3
	return a
}

func AddPointAndReport(member **data.Member, p int) string {

	data.AddPoint(member, p)

	s_string := "<<得点アップサービス>>\n"
	s_string += fmt.Sprintf("%s さんのポイント %d 点アップ\n", (**member).Name, p)
	s_string += "\n"

	return s_string
}

func CreateFriendAndReport(member data.Member, friend_name string) (data.Member, string) {
	friend := data.CreateFriendMember(member, friend_name)

	s_string := fmt.Sprintf("%s さんの紹介で、お友達 %s さんが加わりました", member.Name, friend_name)
	s_string += "\n"

	return friend, s_string
}
