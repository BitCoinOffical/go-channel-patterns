package piepline

import (
	"fmt"
	"math"
)

func main() {
	in := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	out := filterOdd(in)
	out = square(out)
	out = half(out)

	for v := range out {
		fmt.Println(v)
	}

}

func filterOdd(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			val := math.Pow(float64(i), 2)
			out <- int(val)
		}
	}()

	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			val := i / 2
			out <- val
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
