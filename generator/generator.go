package generator

func GenerateWithChannel(start, end int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := start; start <= end; start++ {
			out <- num
		}
	}()

	return out
}
