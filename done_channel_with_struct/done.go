package donechannelwithstruct

import "time"

type Worker struct {
	closeCh       chan struct{}
	closeDoneChan chan struct{}
}

func NewWorker() Worker {
	worker := Worker{
		closeCh:       make(chan struct{}),
		closeDoneChan: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			ticker.Stop()
			close(worker.closeDoneChan)
		}()

		for {
			select {
			case <-worker.closeCh:
				return
			default:
			}

			select {
			case <-worker.closeCh:
				return
			case <-ticker.C:
			}

		}
	}()

	return worker
}

func (w *Worker) ShotDown() {
	close(w.closeCh)
	<-w.closeDoneChan
}
