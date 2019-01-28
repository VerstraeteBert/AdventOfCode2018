package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type f struct {
	fabric [1000][1000]int
	claims []claim
}

type claim struct {
	id    int
	initX int
	initY int
	w     int
	h     int
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
		var c claim
		fmt.Sscanf(sc.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.initX, &c.initY, &c.w, &c.h)
		f.fillFabric(c)
		f.claims = append(f.claims, c)
	}

	goodClaim := f.findGoodClaim()
	fmt.Print(goodClaim)
}

func (f *f) fillFabric(c claim) {
	for y := c.initY; y < c.initY+c.h; y++ {
		for x := c.initX; x < c.initX+c.w; x++ {
			if f.fabric[y][x] == 1 {
				f.fabric[y][x] = -1
			}
			if f.fabric[y][x] == 0 {
				f.fabric[y][x] = 1
			}
		}
	}
}

func (f f) findGoodClaim() int {
	for _, c := range f.claims {
		isGood := true

		for y := c.initY; y < c.initY+c.h; y++ {
			if !isGood {
				break
			}

			for x := c.initX; x < c.initX+c.w; x++ {
				if f.fabric[y][x] != 1 {
					isGood = false
					break
				}
			}
		}

		if isGood {
			return c.id
		}
	}
	// no good claims
	return -1
}
