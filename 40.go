package main

import "fmt"

func main() {
	// m := make(map[string]int)
	// var m = map[string]int{}
	// var m map[string]int
	// m = make(map[string]int)
	// var m map[string]int = map[string]int{"hoge": 10}
	var m = map[string]int{"hoge": 10}

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 48
	v2, ok2 := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)
}
