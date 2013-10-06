package main

import "fmt"

// func fibonacci(x int) int {
//  if x == 0 {
//    return 0
//  } else if x == 1 {
//    return 1
//  } else {
//    return fibonacci(x-2) + fibonacci(x-1)
//  }
// }

// func main() {
//  for i := 0; i < 10; i++ {
//    fmt.Println(fibonacci(i))
//  }
// }

// func fibonacci() func() int {
// 	fib := map[int]int{}
// 	return func() int {
// 		l := len(fib)
// 		if l == 0 {
// 			fib[0] = 0
// 		} else if l == 1 {
// 			fib[1] = 1
// 		} else {
// 			fib[l] = fib[l-2] + fib[l-1]
// 		}
// 		return fib[l]
// 	}
// }

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
