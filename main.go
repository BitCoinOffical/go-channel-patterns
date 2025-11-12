package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/BitCoinOffical/go-channel-patterns/fanout"
)

func main() {

	// start := time.Now()
	// fani := fanin.FanIn(ch1, ch2, ch3)
	// for val := range fani {
	// 	fmt.Println(val)
	// }
	// fmt.Println(time.Since(start))

	// ctx, canel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer canel()
	// ctxfani := fanin.WithContextFanIn(ctx, ch1, ch2, ch3)
	// start := time.Now()
	// for val := range ctxfani {
	// 	fmt.Println(val)
	// }
	// fmt.Println(time.Since(start))

	wg := &sync.WaitGroup{}

	ch1 := make(chan int)

	wg.Go(func() {
		ch1 <- 1
		close(ch1)
	})

	// start := time.Now()
	// chs := fanout.FanOut(ch1, 3)
	// for _, val := range chs {
	// 	fmt.Println(val)
	// }
	// fmt.Println(time.Since(start))

	ctx, canel := context.WithTimeout(context.Background(), 2*time.Second)
	defer canel()
	start := time.Now()
	chs := fanout.WithContextFanOut(ctx, ch1, 3)
	for _, val := range chs {
		wg.Go(func() {
			for v := range val {
				fmt.Println(v)
			}
		})
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
