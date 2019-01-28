package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for i := range words {
		for j := i + 1; j < len(words); j++ {
			common, isClose := compare(words[i], words[j])
			if isClose {
				fmt.Println(common)
			}
		}
	}
}

// checks if words of the same length are only one letter away
// returns the characters in common if they are close
func compare(wordOne, wordTwo string) (common string, isClose bool) {
	if len(wordOne) != len(wordTwo) {
		return "", false
	}

	diffIdx := -1

	for i := range wordOne {
		if wordOne[i] == wordTwo[i] {
			continue
		}
		if diffIdx >= 0 {
			return "", false
		}
		diffIdx = i
	}

	if diffIdx < 0 {
		return "", false
	}

	return wordOne[:diffIdx] + wordOne[diffIdx+1:], true
}
