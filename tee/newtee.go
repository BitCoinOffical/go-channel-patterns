package tee

import (
	"context"
	"sync"
)

type TeeChans struct {
	chans    []chan int
	numChans int
	wg       WaitG
	wgs      []WaitG
}

func NewTee(numChans int, wgSlow WaitG, wgFast WaitG) *TeeChans {
	wgs := make([]WaitG, numChans)
	chans := make([]chan int, numChans)

	for i := range numChans {
		chans[i] = make(chan int)
	}

	for i := range numChans {
		wgs[i] = wgFast
	}

	return &TeeChans{
		chans:    chans,
		numChans: numChans,
		wg:       wgSlow,
		wgs:      wgs,
	}
}

func (t *TeeChans) WitchCtxExecuteNewTee(ctx context.Context, in chan int) []chan int {
	go func() {
		defer func() {
			for i := range t.numChans {
				go func() {
					t.wgs[i].Wait()
					close(t.chans[i])
				}()
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				for i := range t.numChans {
					t.wgs[i].Add(1)
					t.wg.Add(1)
					go func() {
						defer t.wgs[i].Done()
						defer t.wg.Done()

						select {
						case <-ctx.Done():
							return
						case t.chans[i] <- val:

						}

					}()
				}
				t.wg.Wait()
			}
		}
	}()

	return t.chans
}

func (t *TeeChans) ExecuteNewTee(in chan int) []chan int {
	go func() {
		defer func() {
			for i := range t.numChans {
				go func() {
					t.wgs[i].Wait()
					close(t.chans[i])
				}()
			}
		}()

		for val := range in {

			for i := range t.numChans {
				t.wgs[i].Add(1)
				t.wg.Add(1)
				go func() {
					defer t.wgs[i].Done()
					defer t.wg.Done()
					t.chans[i] <- val

				}()
			}
			t.wg.Wait()
		}

	}()

	return t.chans
}
func (t *TeeChans) WithContextTee(ctx context.Context, in chan int) []chan int {
	chans := make([]chan int, t.numChans)
	for i := range t.numChans {
		chans[i] = make(chan int)
	}
	go func() {
		for i := range t.numChans {
			defer close(chans[i])
		}
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				wg := &sync.WaitGroup{}

				for i := range t.numChans {
					wg.Go(func() {
						select {
						case <-ctx.Done():
							return
						case chans[i] <- val:

						}

					})
				}
				wg.Wait()

			}

		}
	}()
	return chans
}

func Tees(in chan int, numchans int) []chan int {

	chans := make([]chan int, numchans)

	for i := range numchans {
		chans[i] = make(chan int)
	}

	go func() {

		for i := range numchans {
			defer close(chans[i])
		}

		for v := range in {
			for i := range numchans {

				chans[i] <- v
			}
		}

	}()

	return chans
}
