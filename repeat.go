package handy

func Repeat(times int, fn func(offset int)) {
	for i := 0; i < times; i++ {
		fn(i)
	}
}
