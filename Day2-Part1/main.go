package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
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
		go func (s string) {
			result := getDupAndTripLetters(s)
			res.add(result)
			wg.Done()
		} (scanner.Text())
	}

	wg.Wait()

	fmt.Printf("The checksum (%d * %d) is %d", res.double, res.triple, res.double * res.triple)
}

type numberOcc struct {
	double int
	triple int
	mu sync.Mutex
}

func (occ *numberOcc) add (occ2 numberOcc) {
	occ.mu.Lock()
	occ.double += occ2.double
	occ.triple += occ2.triple
	occ.mu.Unlock()
}

// Gets the letters that occur twice and three times
// Returns the both of these through a channel as a numberOcc struct
func getDupAndTripLetters(s string) (result numberOcc) {
	letterFreqs := make([]int, 26)

	for _, char := range s {
		letterFreqs[char - 97]++
	}

	for _, occ := range letterFreqs {
		if occ == 2 {
			result.double = 1
		} else if occ == 3 {
			result.triple = 1
		}
	}

	return
}
