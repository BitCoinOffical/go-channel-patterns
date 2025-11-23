package promis

type Result struct {
	val int
	err error
}

type Promise struct {
	resultCh chan Result
}

func NewPromis(asyncFn func() (int, error)) Promise {
	promise := Promise{
		resultCh: make(chan Result),
	}

	go func() {
		defer close(promise.resultCh)

		val, err := asyncFn()
		promise.resultCh <- Result{val: val, err: err}
	}()

	return promise
}

func (p *Promise) Then(succes func(int), errorFn func(error)) {
	go func() {
		result := <-p.resultCh
		if result.err != nil {
			errorFn(result.err)
		}
		succes(result.val)
	}()
}
