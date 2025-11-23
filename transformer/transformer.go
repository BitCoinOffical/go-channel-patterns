package transformer

func Transform(in chan int, f func(int) int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- f(num)
		}
	}()

	return out
}
