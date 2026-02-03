package parallelforloop

import (
	"fmt"
	"sync"
	"time"
)

type empty struct{}

const data_size = 10
const semaphore_limit = 3

func main() {
	data := make([]int, 00, data_size)

	for i := 0; i < data_size; i++ {
		data = append(data, i+1)
	}

	res := make([]int, data_size)
	sem := make(chan empty, semaphore_limit)
	wg := &sync.WaitGroup{}
	fmt.Println("before", data)

	start := time.Now()

	for i, xi := range data {
		wg.Add(1)
		go func(i, xi int) {
			defer wg.Done()
			sem <- empty{}
			res[i] = calc(xi)
			<-sem
		}(i, xi)
	}

	wg.Wait()
	fmt.Println("After:", res)
	fmt.Println("Elapsed:", time.Since(start))
}

func calc(val int) int {
	fmt.Println("calc", val)
	time.Sleep(500 * time.Millisecond)
	return val * 2
}
