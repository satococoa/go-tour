package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  p = Vertex{1, 2}
  q = &Vertex{1, 2}
  r = Vertex{X: 1}
  s = Vertex{}
  t = Vertex{Y: 2}
  // u = Vertex{, 3} // error
  // v = Vertex{2} // error
)

func main() {
  fmt.Println(p, q, r, s, t)
}