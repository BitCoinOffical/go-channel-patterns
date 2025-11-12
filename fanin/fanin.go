package fanin

import (
	"context"
	"sync"
)

func WithContextFanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	wg := &sync.WaitGroup{}
	ch := make(chan T)

	for _, channel := range chans {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return

				case val, ok := <-channel:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return

					case ch <- val:
					}

				}
			}
		})
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func FanIn[T any](chans ...<-chan T) <-chan T {
	wg := &sync.WaitGroup{}
	ch := make(chan T)
	wg.Add(len(chans))

	for _, channel := range chans {

		go func() {
			defer wg.Done()
			for val := range channel {
				ch <- val
			}
		}()

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
