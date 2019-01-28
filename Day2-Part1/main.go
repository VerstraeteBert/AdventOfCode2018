package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	fmt.Println(concurrent())
	fmt.Println(nonConcurrent())
}

func concurrent() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := new(numberOcc)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(s string) {
			res.add(getDupAndTripLetters(s))
			wg.Done()
		}(scanner.Text())
	}

	wg.Wait()

	return res.twos * res.threes
}

type numberOcc struct {
	twos   int
	threes int
	mu     sync.Mutex
}

func (occ *numberOcc) add(gotTwo, gotThree bool) {
	occ.mu.Lock()
	if gotTwo {
		occ.twos++
	}
	if gotThree {
		occ.threes++
	}
	occ.mu.Unlock()
}

func nonConcurrent() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numTwos, numThrees int

	for scanner.Scan() {
		gotTwo, gotThree := getDupAndTripLetters(scanner.Text())
		if gotTwo {
			numTwos++
		}
		if gotThree {
			numThrees++
		}
	}

	return numTwos * numThrees
}

// checks if any letter occurs exactly twice or three times in a string
func getDupAndTripLetters(s string) (gotTwo, gotThree bool) {
	letterFreqs := make(map[rune]int)

	for _, char := range s {
		letterFreqs[char]++
	}

	for _, occ := range letterFreqs {
		if occ == 2 {
			gotTwo = true
		} else if occ == 3 {
			gotThree = true
		}
	}

	return
}
