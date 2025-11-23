package semaphor

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(tickNum int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, tickNum),
	}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.tickets
}
