package test

import (
  "testing"

  "go-optimized-web-crawler/set"
)

func TestAdd(t *testing.T) {
  defer dontPanic(t)
  s := set.New()
  s.Add(1)
  s.Add(2)
  s.Add(3)
  s.Add(4)
}

func TestAddAsync(t *testing.T) {
  defer dontPanic(t)
  s := set.New()
  go s.Add(1)
  go s.Add(2)
  go s.Add(3)
  go s.Add(4)
}

func TestDelAsync(t *testing.T) {
  defer dontPanic(t)
  s := set.New()
  s.Add(1)
  s.Add(2)
  s.Add(3)
  s.Add(4)
  go s.Del(1)
  go s.Del(2)
  go s.Del(3)
  go s.Del(4)
}

func TestAddAndDel(t *testing.T) {
  defer dontPanic(t)
  s := set.New()
  s.Add(1)
  s.Add(3)
  compareValue(t, s.Has(3), true)
  s.Del(3)
  compareValue(t, s.Has(3), false)
  s.Add(4)
  s.Del(1)
  s.Add(2)
  compareValue(t, s.Has(5), false)
  s.Del(3)
  s.Add(4)
  compareValue(t, s.Has(4), true)
}
