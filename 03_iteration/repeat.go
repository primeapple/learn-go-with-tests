package iteration

func Repeat(s string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}
