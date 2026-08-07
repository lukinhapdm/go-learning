package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Initialize an empty map to hold the word counts
	counts := make(map[string]int)
	
	// Split the string into an array of words
	words := strings.Fields(s)
	
	// Iterate over the words and increment their count in the map
	for _, word := range words {
		counts[word]++
	}
	
	return counts
}

func main() {
	wc.Test(WordCount)
}