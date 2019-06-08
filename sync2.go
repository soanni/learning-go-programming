package main

import (
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.Mutex
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.mutex.Unlock()
		<-s.stpCh
	}()
}

func (s *Service) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
}

func main() {
	svc := &Service{}
	svc.Start()
	time.Sleep(time.Second)
	svc.Stop()
}
