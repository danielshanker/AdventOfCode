package main

import (
	"flag"
	"math"
	. "utils"
)

func main() {
	debug := false
	test := flag.Bool("t", debug, "use sample")
	flag.Parse()

	Start(test, 8, part1, part2, 14, 34)

}

type ant struct {
	orig  string
	found bool
	coord
}

type coord struct {
	x int
	y int
}

func part1(lines []string) int {
	answer := 0

	ants := map[coord]ant{}

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			if char != "." {
				c := coord{x: j, y: i}
				a := ant{
					coord: c,
					orig:  char,
				}
				ants[c] = a
			}
		}
	}

	for c, a := range ants {
		if a.orig == "" {
			continue
		}
		for newC, newA := range ants {
			if newA.orig != a.orig || newC == c {
				continue
			}
			distX := newC.x - c.x
			distY := newC.y - c.y
			newX := c.x - distX
			newY := c.y - distY
			if newX < 0 || newX >= len(lines[0]) || newY < 0 || newY >= len(lines) {
				continue
			}
			xy := coord{x: newX, y: newY}

			newA := ants[xy]
			newA.found = true
			ants[xy] = newA
		}
	}

	for _, a := range ants {
		if a.found {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	ants := map[coord]ant{}

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			if char != "." {
				c := coord{x: j, y: i}
				a := ant{
					coord: c,
					orig:  char,
					found: true,
				}

				ants[c] = a
			}
		}
	}

	maxL := int(math.Max(float64(len(lines)), float64(len(lines[0]))))

	for c, a := range ants {
		if a.orig == "" {
			continue
		}
		for newC, newA := range ants {
			if newA.orig != a.orig || newC == c {
				continue
			}
			newX := c.x
			newY := c.y
			distX := newC.x - c.x
			distY := newC.y - c.y
			for i := 0; i < maxL; i++ {
				newX = newX - distX
				newY = newY - distY
				if newX < 0 || newX >= len(lines[0]) || newY < 0 || newY >= len(lines) {
					continue
				}
				xy := coord{x: newX, y: newY}

				newAnt := ants[xy]
				newAnt.found = true
				ants[xy] = newAnt
			}
		}
	}

	for _, a := range ants {
		if a.found {
			answer++
		}
	}

	return answer
}
