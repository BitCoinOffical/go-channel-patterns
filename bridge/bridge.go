package bridge

func Bridge(in chan chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for ch := range in {
			for val := range ch {
				out <- val
			}
		}
	}()
	return out
}
