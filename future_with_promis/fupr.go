package futurewithpromis

type Future struct {
	res chan int
}

func NewFuture(res chan int) Future {
	return Future{
		res: res,
	}
}

func (f *Future) Get() int {
	return <-f.res
}

type Promise struct {
	res chan int
}

func NewPromis() Promise {
	return Promise{
		res: make(chan int),
	}
}

func (p *Promise) Set(val int) {
	p.res <- val
	close(p.res)
}
func (p *Promise) GetFuture() Future {
	return NewFuture(p.res)
}
