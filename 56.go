package main
import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return x, ErrNegativeSqrt(x)
  }
  minimum_delta := 1 / 10000000000000000000.0
  z := 1.0
  for i := 0; i < 10; i++ {
    // fmt.Printf("%d\n", i)
    new_z := z - ((z * z) - x) / (2 * z)
    if new_z > z && new_z - z < minimum_delta ||
    new_z < z && z - new_z < minimum_delta {
      return new_z, nil
    } else {
      z = new_z
    }
  }
  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
