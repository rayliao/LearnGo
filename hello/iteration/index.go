package iteration

// 如何声明默认值？
func Repeat(r string, t int) string {
	var repeated string
	for i := 0; i < t; i++ {
		repeated += r
	}
	return repeated
}
