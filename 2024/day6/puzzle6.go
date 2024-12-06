package main

import (
	"flag"
	. "utils"

	"golang.org/x/exp/slices"
)

func main() {
	debug := false
	test := flag.Bool("t", debug, "use sample")
	flag.Parse()

	Start(test, 6, part1, part2, 41, 6)

}

type spot struct {
	visited     bool
	obstructed  bool
	visitedDirs []int
}

type coord struct {
	x int
	y int
}

const (
	up    = 1
	right = 2
	down  = 3
	left  = 4
)

var originalVisitedSpots = map[coord]bool{}

func part1(lines []string) int {
	answer := 0
	spots := map[coord]spot{}

	pos := coord{0, 0}
	dir := up

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			c := coord{x: j, y: i}
			if char == "#" {
				spots[c] = spot{obstructed: true}
			}
			if char == "^" {
				spots[c] = spot{visited: true}
				pos.x = j
				pos.y = i
			}
		}
	}

	for true {
		if pos.x < 0 || pos.x > len(lines[0])-1 || pos.y < 0 || pos.y > len(lines) {
			break
		}
		newPos := pos
		if dir == up {
			newPos = coord{pos.x, pos.y - 1}
			if spots[newPos].obstructed {
				dir = right
				continue
			}
		} else if dir == right {
			newPos = coord{pos.x + 1, pos.y}
			if spots[newPos].obstructed {
				dir = down
				continue
			}
		} else if dir == down {
			newPos = coord{pos.x, pos.y + 1}
			if spots[newPos].obstructed {
				dir = left
				continue
			}
		} else if dir == left {
			newPos = coord{pos.x - 1, pos.y}
			if spots[newPos].obstructed {
				dir = up
				continue
			}
		}

		a := spots[newPos]
		a.visited = true
		spots[newPos] = a
		originalVisitedSpots[newPos] = true
		pos = newPos
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			c := coord{
				x: j,
				y: i,
			}

			if spots[c].visited {
				answer++
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	originalSpots := map[coord]spot{}

	pos := coord{0, 0}

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			c := coord{x: j, y: i}
			if char == "#" {
				originalSpots[c] = spot{obstructed: true}
			}
			if char == "^" {
				originalSpots[c] = spot{visited: true}
				pos.x = j
				pos.y = i
			}
		}
	}
	originalPos := pos

	for i, line := range lines {
		for j := range line {
			dir := up
			pos = originalPos
			newObs := coord{
				x: j,
				y: i,
			}
			if originalSpots[newObs].obstructed || newObs == pos || !originalVisitedSpots[newObs] {
				continue
			}
			spots := copyMap(originalSpots)
			spots[newObs] = spot{obstructed: true}
			for true {
				if pos.x < 0 || pos.x > len(lines[0])-1 || pos.y < 0 || pos.y > len(lines) {
					break
				}
				newPos := pos
				if dir == up {
					newPos = coord{pos.x, pos.y - 1}
					if spots[newPos].obstructed {
						dir = right
						continue
					}
				} else if dir == right {
					newPos = coord{pos.x + 1, pos.y}
					if spots[newPos].obstructed {
						dir = down
						continue
					}
				} else if dir == down {
					newPos = coord{pos.x, pos.y + 1}
					if spots[newPos].obstructed {
						dir = left
						continue
					}
				} else if dir == left {
					newPos = coord{pos.x - 1, pos.y}
					if spots[newPos].obstructed {
						dir = up
						continue
					}
				}

				if spots[newPos].visited && slices.Contains(spots[newPos].visitedDirs, dir) {
					answer++
					break
				}
				a := spots[newPos]
				a.visited = true
				a.visitedDirs = append(a.visitedDirs, dir)
				spots[newPos] = a
				pos = newPos
			}
		}
	}

	return answer
}

func copyMap(orig map[coord]spot) map[coord]spot {
	copy := make(map[coord]spot)
	for k, v := range orig {
		copy[k] = v
	}
	return copy
}
