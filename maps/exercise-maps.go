package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	// Split the string into words
	words := strings.Fields(s)

	for _, word := range words {
		// Check if the word exists in the map
		_, ok := res[word]
		if ok {
			res[word]++
		} else {
			res[word] = 1
		}
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
