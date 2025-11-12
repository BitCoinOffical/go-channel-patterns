package tee

import (
	"context"
	"sync"
)

func Tee(ctx context.Context, in chan int, numChans int) []chan int {
	chans := make([]chan int, numChans)
	for i := range numChans {
		chans[i] = make(chan int)
	}
	go func() {
		for i := range numChans {
			defer close(chans[i])
		}
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				wg := &sync.WaitGroup{}

				for i := range numChans {
					wg.Go(func() {
						select {
						case <-ctx.Done():
							return
						case chans[i] <- val:

						}

					})
				}
				wg.Wait()

			}

		}
	}()
	return chans
}
