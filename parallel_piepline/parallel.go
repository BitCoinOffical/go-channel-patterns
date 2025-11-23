package parallelpiepline

import "sync"

func Split(value chan int, numChans int) []chan int {
	out := make([]chan int, numChans)

	for i := range numChans {
		out[i] = make(chan int)
	}
	go func() {
		defer func() {
			for i := range numChans {
				close(out[i])
			}
		}()
		for v := range value {
			for i := range numChans {
				out[i] <- v
			}
		}
	}()

	return out
}

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

func PiePline(in chan int, n int) chan int {
	wg := &sync.WaitGroup{}
	out := make(chan int)

	for range n {
		wg.Go(func() {
			for i := range in {
				out <- i
			}
		})
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
