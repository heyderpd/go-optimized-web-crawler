package tailRecursion

import (
  "go-optimized-web-crawler/utils"
)

type TailRecursion struct {
  tasks []interface{}
}

func New() *TailRecursion {
  s := TailRecursion{}
  return &s
}

func (s *TailRecursion) AppendTask(task interface{}) {
  s.tasks = append(s.tasks, task)
  // log.Println("AppendTask.length", len(s.tasks))
}

func (s *TailRecursion) getNextTask() interface{} {
  length := len(s.tasks)
  // log.Println("getNextTask.length", length)
  if length == 0 {
    return nil
  }
  task := s.tasks[0]
  s.tasks = s.tasks[1:]
  return task
}

func (s *TailRecursion) Recursion(firstTask interface{}, hanlder utils.UnknownHandler) {
  s.tasks = make([]interface{}, 0)
  s.AppendTask(firstTask)
  for t := s.getNextTask(); t != nil; t = s.getNextTask() {
    // log.Println("t", t)
    hanlder(t)
  }
}
