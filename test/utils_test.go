package test

import (
  "testing"
  "log"
  "reflect"
  "time"
)

func needPanic(t *testing.T) {
  r := recover()
	if r == nil {
    t.Errorf("Expected to panic")
  }
}

func dontPanic(t *testing.T) {
  r := recover()
	if r != nil {
    t.Errorf("Unexpected to panic")
  }
}

func compareValue(t *testing.T, result interface{}, expected interface{}) {
  a := reflect.ValueOf(result)
  b := reflect.ValueOf(expected)
  if a != b {
    log.Println("result", result)
    log.Println("expected", expected)
    t.Errorf("Value don't match")
  }
}

func getNow() float64 {
	return float64(time.Now().UnixNano())
}
