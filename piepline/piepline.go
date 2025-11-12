package piepline

import (
	tasker "github.com/BitCoinOffical/go-channel-patterns/pkg"
)

func Pipeline[T any](in <-chan T) chan T {
	out := make(chan T)
	go func() {
		for v := range in {
			out <- tasker.Task(v)
		}
		close(out)
	}()
	return out
}
