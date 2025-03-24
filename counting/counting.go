package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Go is an open source programming language that makes it easy to build 
    simple, reliable, and efficient software. Go was designed at Google in 2007 
    by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar 
    to C, but with memory safety, garbage collection, structural typing, and 
    CSP-style concurrency.`

	// Count word frequencies
	words := strings.Fields(strings.ToLower(text))
	frequencies := make(map[string]int)

	for _, word := range words {
		// Remove punctuation (simplified approach)
		word = strings.Trim(word, ".,;:!?()[]{}\"'")
		frequencies[word]++
	}

	fmt.Println("Word frequencies:")
	for word, count := range frequencies {
		if count > 1 {
			fmt.Printf("%-12s: %d\n", word, count)
		}
	}
}
