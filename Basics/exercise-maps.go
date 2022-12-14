package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		m[word] += 1
	}
	return m //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
