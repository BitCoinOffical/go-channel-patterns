package filter

func Filter(in chan int, filterf func(int) bool) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			if filterf(num) {
				out <- num
			}
		}
	}()

	return out
}
