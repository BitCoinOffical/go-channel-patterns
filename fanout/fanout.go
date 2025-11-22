package fanout

import (
	"context"
	"sync"
)

func WithContextFanOut[T any](ctx context.Context, in chan T, numChans int) []chan T {
	wg := &sync.WaitGroup{}
	chans := make([]chan T, numChans)

	for i := range numChans {
		chans[i] = make(chan T)
	}

	wg.Go(func() {

		defer func() {
			for _, c := range chans {
				close(c)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				for _, c := range chans {
					select {
					case <-ctx.Done():
						return
					case c <- val:
					}
				}
			}
		}
	})

	go func() {
		wg.Wait()

	}()

	return chans
}

func FanOut[T any](in chan T, numChans int) []chan T {
	chans := make([]chan T, numChans)

	for i := range numChans {
		chans[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, с := range chans {
				close(с)
			}
		}()

		for v := range in {
			for _, i := range chans {
				i <- v
			}
		}
	}()

	return chans
}
