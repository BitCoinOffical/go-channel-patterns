package generator

import (
	"fmt"
	"sync"
	"time"
)

func makeGenerator(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	out := make(chan int, 1)
	i := 0
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(out)
				fmt.Print("done")
				return
			default:
				time.Sleep(500 * time.Millisecond)
				out <- i
				i++
			}
		}
	}()

	return out
}

func main() {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch := makeGenerator(done, wg)

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value:", v)
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
