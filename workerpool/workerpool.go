package workerpool

import "sync"

func WorkerPool[T any](in <-chan T, workerCount int, f func(T) T) chan T {
	wg := &sync.WaitGroup{}
	out := make(chan T)

	for range workerCount {
		wg.Go(func() {
			for v := range in {
				out <- f(v)
			}
		})
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
