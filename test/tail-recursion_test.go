package test

import (
  "testing"
  "runtime"
  "log"

  "go-optimized-web-crawler/tail-recursion"
)

type Node struct {
  val int
  nodes []Node
}

var (
  testData = Node{
    val: 0,
    nodes: []Node{
      Node{
        val: 1,
        nodes: []Node{
          Node{
            val: 4,
            nodes: []Node{},
          },
          Node{
            val: 5,
            nodes: []Node{
              Node{
                val: 7,
                nodes: []Node{},
              },
              Node{
                val: 8,
                nodes: []Node{},
              },
            },
          },
        },
      },
      Node{
        val: 2,
        nodes: []Node{},
      },
      Node{
        val: 3,
        nodes: []Node{
          Node{
            val: 6,
            nodes: []Node{},
          },
        },
      },
    },
  }
)

func TestTailRecursion(t *testing.T) {
  defer dontPanic(t)
  tr := tailRecursion.New()
  var rmax int = 0
  order := make([]int, 0)
  tr.Recursion(testData, func (i interface{}) {
    r := runtime.NumGoroutine()
    if r > rmax {
      rmax = r
    }
    node := (i).(Node)
    order = append(order, node.val)
    for _, n := range node.nodes {
      tr.AppendTask(n)
    }
  })
  expected := []int{0,1,2,3,4,5,6,7,8}
  for k, e := range expected {
    v := order[k]
    if e != v {
      log.Println("value", v)
      log.Println("expected", e)
      t.Errorf("Fail on order tail recursion")
      return
    }
  }
  if rmax > 3  {
    t.Errorf("Behavior not expected")
  }
}
