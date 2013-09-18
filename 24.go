package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  minimum_delta := 1 / 10000000000000000000.0
  z := 1.0
  for i := 0; i < 10; i++ {
    fmt.Printf("%d\n", i)
    new_z := z - ((z * z) - x) / (2 * z)
    if new_z > z && new_z - z < minimum_delta ||
    new_z < z && z - new_z < minimum_delta {
      return new_z
    } else {
      z = new_z
    }
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}