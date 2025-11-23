package future

type Future struct {
	res chan int
}

func NewFuture(action func() int) Future {
	future := Future{
		res: make(chan int),
	}

	go func() {
		defer close(future.res)
		future.res <- action()
	}()

	return future
}

func (f *Future) Get() int {
	return <-f.res
}
