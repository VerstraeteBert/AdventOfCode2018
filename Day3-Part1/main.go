package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type f struct {
	fabric     [1000][1000]int
	numClashes int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	f := new(f)

	for sc.Scan() {
		var id, initX, initY, w, h int
		fmt.Sscanf(sc.Text(), "#%d @ %d,%d: %dx%d", &id, &initX, &initY, &w, &h)
		f.fillFabric(initX, initY, w, h)
	}

	fmt.Printf("There are %d overlaps", f.numClashes)
}

func (f *f) fillFabric(initX, initY, w, h int) {
	for y := initY; y < initY+h; y++ {
		for x := initX; x < initX+w; x++ {
			if f.fabric[y][x] == 1 {
				f.numClashes++
				f.fabric[y][x] = -1
			}
			if f.fabric[y][x] == 0 {
				f.fabric[y][x] = 1
			}
		}
	}
}
