package errgroup

import "sync"

type ErrGroup struct {
	err  error
	wg   sync.WaitGroup
	once sync.Once

	done chan struct{}
}

func NewErrGroup() (*ErrGroup, chan struct{}) {
	dpne := make(chan struct{})
	return &ErrGroup{
		done: dpne,
	}, dpne
}

func (eg *ErrGroup) Go(Task func() error) {
	eg.wg.Add(1)
	go func() {
		defer eg.wg.Done()

		select {
		case <-eg.done:
			return
		default:
			if err := Task(); err != nil {
				eg.once.Do(func() {
					eg.err = err
					close(eg.done)
				})
			}
		}
	}()

}

func (eg *ErrGroup) Wait() error {
	eg.wg.Wait()
	return eg.err
}
