package fanout

import (
	"context"

	"github.com/BitCoinOffical/go-channel-patterns/piepline"
)

func WithContextFanOut[T any](ctx context.Context, in <-chan T, numChans int) []chan T {
	chans := make([]chan T, numChans)
	go func() {
		for i := range numChans {
			select {
			case <-ctx.Done():
				return
			default:
				chans[i] = piepline.Pipeline(in)
			}

		}
	}()
	return chans
}

func FanOut[T any](in <-chan T, numChans int) []chan T {
	chans := make([]chan T, numChans)
	for i := range numChans {
		chans[i] = piepline.Pipeline(in)
	}
	return chans
}
