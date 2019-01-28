package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		initX, initY, w, h := parseDimensions(sc.Text())
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

func parseDimensions(s string) (x, y, w, h int) {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '#' || r == '@' || r == ',' || r == 'x' || r == ':' || r == ' '
	})

	x, _ = strconv.Atoi(parts[1])
	y, _ = strconv.Atoi(parts[2])
	w, _ = strconv.Atoi(parts[3])
	h, _ = strconv.Atoi(parts[4])

	return
}
