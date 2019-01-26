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

	result := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var n int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}

		result += n
	}

	fmt.Printf("The result is %d", result)
}
