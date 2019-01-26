package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nums []int
	for scanner.Scan() {
		var num int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &num)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	currSum := 0
	freqMemo := make(map[int]bool)
	freqMemo[currSum] = true

	for {
		for _, num := range nums {
			currSum += num

			if _, present := freqMemo[currSum]; present {
				fmt.Printf("The first frequency reached twice is %d", currSum)
				return
			}

			freqMemo[currSum] = true
		}
	}
}