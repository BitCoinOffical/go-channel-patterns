package singlflight

import (
	"sync"
)

type call struct {
	err error
	val interface {
	}
	done chan struct {
	}
}

type SinglFlight struct {
	mutex sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SinglFlight {
	return &SinglFlight{
		calls: make(map[string]*call),
	}
}

func (s *SinglFlight) Wait(call *call) (any, error) {
	<-call.done
	return call.val, call.err
}

func (s *SinglFlight) Do(key string, action func() (interface{}, error)) (any, error) {
	s.mutex.Lock()
	if call, found := s.calls[key]; found {
		s.mutex.Unlock()
		return s.Wait(call)
	}

	call := &call{
		done: make(chan struct{}),
	}

	s.calls[key] = call
	s.mutex.Unlock()

	go func() {
		defer func() {
			s.mutex.Lock()
			close(call.done)
			delete(s.calls, key)
			s.mutex.Unlock()
		}()
		call.val, call.err = action()
	}()
	return s.Wait(call)

}
