package piepline

func Generate(values ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range values {
			out <- i
		}
	}()

	return out
}

func PiePline(in chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			out <- f(i)
		}
	}()

	return out
}
