package main

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	sync.RWMutex
	cache map[int]string
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	s.cache = make(map[int]string)
	go func() {
		s.Lock()
		s.started = true
		s.cache[1] = "Entry 1"
		s.Unlock()
		<-s.stpCh
	}()
}

func (s *Service) Serve(id int) {
	s.RLock()
	msg := s.cache[id]
	s.RUnlock()
	if msg != "" {
		fmt.Println(msg)
	} else {
		fmt.Println("Goodbye!")
	}
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
	svc.Serve(1)
	svc.Stop()
}
