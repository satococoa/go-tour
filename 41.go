package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, v := range strings.Fields(s) {
		i, ok := result[v]
		if ok {
			result[v] = i + 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func WordCount2(s string) (x map[string]int) {
	x = make(map[string]int)
	for _, v := range strings.Fields(s) {
		i, ok := x[v]
		if ok {
			x[v] = i + 1
		} else {
			x[v] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
	wc.Test(WordCount2)
}
