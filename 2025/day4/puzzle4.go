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

type cell struct {
	coord    Coord
	hasPaper bool
}

func part1(lines []string) int {
	answer := 0

	grid := make(map[Coord]cell)

	for i, line := range lines {
		for j, r := range line {
			coord := Coord{
				X: j,
				Y: i,
			}
			c := cell{
				coord:    coord,
				hasPaper: string(r) == "@",
			}
			grid[coord] = c
		}
	}

	for _, cell := range grid {
		if !cell.hasPaper {
			continue
		}
		if checkNeightbours(grid, cell.coord) {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	grid := make(map[Coord]cell)

	for i, line := range lines {
		for j, r := range line {
			coord := Coord{
				X: j,
				Y: i,
			}
			c := cell{
				coord:    coord,
				hasPaper: string(r) == "@",
			}
			grid[coord] = c
		}
	}

	for {
		changes := []Coord{}
		for _, cell := range grid {
			if !cell.hasPaper {
				continue
			}
			if checkNeightbours(grid, cell.coord) {
				answer++
				changes = append(changes, cell.coord)
			}
		}
		if len(changes) > 0 {
			for _, c := range changes {
				cell := grid[c]
				cell.hasPaper = false
				grid[c] = cell
			}
		} else {
			break
		}
	}

	return answer
}

func checkNeightbours(grid map[Coord]cell, coord Coord) bool {
	paperNeighbours := 0
	neighbours := []Coord{
		{X: coord.X, Y: coord.Y - 1},
		{X: coord.X + 1, Y: coord.Y - 1},
		{X: coord.X - 1, Y: coord.Y - 1},
		{X: coord.X, Y: coord.Y + 1},
		{X: coord.X + 1, Y: coord.Y + 1},
		{X: coord.X - 1, Y: coord.Y + 1},
		{X: coord.X + 1, Y: coord.Y},
		{X: coord.X - 1, Y: coord.Y},
	}
	for _, n := range neighbours {
		if grid[n].hasPaper {
			paperNeighbours++
		}
		if paperNeighbours >= 4 {
			return false
		}
	}

	return paperNeighbours < 4

}
