package main

import (
	"flag"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 4, part1, part2, 13, 43)

}

func part1(lines []string) int {
	answer := 0

	grid := make(map[Coord]bool)

	for i, line := range lines {
		for j, r := range line {
			if r == '@' {
				coord := Coord{
					X: j,
					Y: i,
				}
				grid[coord] = true
			}
		}
	}

	for coord := range grid {
		if checkNeightbours(grid, coord) {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	grid := make(map[Coord]bool)

	for i, line := range lines {
		for j, r := range line {
			if r == '@' {
				coord := Coord{
					X: j,
					Y: i,
				}
				grid[coord] = true
			}
		}
	}

	for {
		changes := []Coord{}
		for coord := range grid {
			if checkNeightbours(grid, coord) {
				answer++
				changes = append(changes, coord)
				delete(grid, coord)
			}
		}
		if len(changes) == 0 {
			break
		}
	}

	return answer
}

var neighbourOffset = []Coord{
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: -1, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: -1, Y: 1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func checkNeightbours(grid map[Coord]bool, coord Coord) bool {
	paperNeighbours := 0
	for _, n := range neighbourOffset {
		nb := Coord{
			X: coord.X + n.X,
			Y: coord.Y + n.Y,
		}
		if grid[nb] {
			paperNeighbours++
			if paperNeighbours >= 4 {
				return false
			}
		}
	}

	return true

}
