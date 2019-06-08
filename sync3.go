package main

import (
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	sync.Mutex
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.Lock()
		s.started = true
		s.Unlock()
		<-s.stpCh
	}()
}

func (s *Service) Stop() {
	s.Lock()
	defer s.Unlock()
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
