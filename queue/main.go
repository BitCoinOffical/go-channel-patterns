package queue

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan struct{}, 2)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		Queue("hi", queue, wg)
	}
	wg.Wait()
}

func Queue(playload string, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Println(playload)
		<-queue
	}()

}
