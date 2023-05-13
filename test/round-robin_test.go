package test

import (
  "testing"
  "sync"
  "time"
  "runtime"

  "go-optimized-web-crawler/round-robin"
)

func TestRoundRobin(t *testing.T) {
  defer dontPanic(t)
  var test = func(routines int) (float64, int) {
    waitGroup := sync.WaitGroup{}
    rr := roundRobin.New(routines)
    start := getNow()
    var rmax int = 0
    for i := 0; i < 50; i++ {
      waitGroup.Add(1)
      rr.Promise(func (x interface{}) {
        r := runtime.NumGoroutine()
        if r > rmax {
          rmax = r
        }
        time.Sleep(10 * time.Millisecond)
        waitGroup.Done()
      })
    }
    waitGroup.Wait()
    end := getNow()
    return end - start, rmax
  }
  timeWith2Parallel, rmaxWith2 := test(2)
  timeWith8Parallel, rmaxWith8 := test(8)
  if rmaxWith8 - 2 >= rmaxWith2 && rmaxWith2 >= rmaxWith8 + 2  {
    t.Errorf("Behavior not expected")
  }
  if timeWith8Parallel >= timeWith2Parallel {
    t.Errorf("Fail on control parallel routines")
  }
}
