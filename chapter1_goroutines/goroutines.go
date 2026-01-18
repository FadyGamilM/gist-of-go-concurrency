package main

import (
	"strings"
	"sync"
	"unicode"
)

func main() {
	str := "10 2020"
	var wg sync.WaitGroup

	words := strings.Fields(str)

	for _, word := range words {
		wg.Go(func() {
			countDigits(word)
		})
	}

	wg.Wait()
	println("DONE!")
}

// countDigits returns the number of digits in a string.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}
