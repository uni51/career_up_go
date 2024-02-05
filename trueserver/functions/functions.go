package functions

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
