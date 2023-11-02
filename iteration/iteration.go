package iteration

func Repeat(s string, x int) string {
	result := ""
	for i := 0; i < x; i++ {
		result += s
	}
	return result
}

func BadLoopVarUsage() int {
	values := [3]int{1, 2, 3}
	pointers := make([]*int, 3)

	for i, v := range values {
		pointers[i] = &v
	}
	sum := 0
	for _, v := range pointers {
		sum += *v
	}
	return sum
}
