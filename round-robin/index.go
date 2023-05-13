package roundRobin

import (
  "sync"

  "go-optimized-web-crawler/utils"
)

type RoundRobin struct {
  next int
  routines int
  semaphores []*sync.Mutex
}

func New(routines int) *RoundRobin {
  if routines < 0 {
    panic("routines must be great than 0")
  }
  if routines > 20 {
    panic("routines must be less than 20")
  }
  s := RoundRobin{
    next: 0,
    routines: routines,
    semaphores: make([]*sync.Mutex, routines),
  }
  for i := 0; i < routines; i++ {
    s.semaphores[i] = new(sync.Mutex)
  }
  return &s
}

func (s *RoundRobin) getNextMutex() *sync.Mutex {
  s.next++
  if s.next >= s.routines {
    s.next = 0
  }
  return s.semaphores[s.next]
}

func (s *RoundRobin) Promise(hanlder utils.UnknownHandler) {
  go func() {
    mutext := s.getNextMutex()
    mutext.Lock()
    hanlder(nil)
    mutext.Unlock()
  }()
}
