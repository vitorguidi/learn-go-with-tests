package iteration

func Repeat(s string, x int) string {
	result := ""
	for i := 0; i < x; i++ {
		result += s
	}
	return result
}
