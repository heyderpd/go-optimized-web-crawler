package set

import (
  "sync"
)

type InterfaceSet map[interface{}]bool

type Set struct {
  semaphore sync.Mutex
  set InterfaceSet
}

func New() *Set {
  s := Set{
    set: InterfaceSet{},
    semaphore: sync.Mutex{},
  }
  return &s
}

func (s *Set) _has(value interface{}) bool {
  _, exist := s.set[value]
  return exist
}

func (s *Set) Has(value interface{}) bool {
  s.semaphore.Lock()
  defer s.semaphore.Unlock()
  return s._has(value)
}

func (s *Set) Add(value interface{}) {
  s.semaphore.Lock()
  defer s.semaphore.Unlock()
  s.set[value] = true
}

func (s *Set) Del(value interface{}) {
  s.semaphore.Lock()
  defer s.semaphore.Unlock()
  has := s._has(value)
  if has {
    delete(s.set, value)
  }
}

func (s *Set) ToList() []interface{} {
  s.semaphore.Lock()
  defer s.semaphore.Unlock()
  list := make([]interface{}, 0)
	for value, _ := range s.set {
    list = append(list, value)
	}
  return list
}
