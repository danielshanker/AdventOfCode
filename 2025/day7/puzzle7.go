package main

import (
	"flag"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 7, part1, part2, 21, 40)

}

type cell struct {
	coord    Coord
	splitter bool
	beam     bool
	downs    int
}

func part1(lines []string) int {
	answer := 0

	beamRow := make(map[int]bool)
	grid := map[Coord]cell{}

	for i, line := range lines {
		for j, r := range line {
			coord := Coord{X: j, Y: i}
			curCell := cell{coord: coord}
			if r == 'S' {
				curCell.beam = true
				grid[coord] = curCell
				beamRow[j] = true
				continue
			}

			if grid[coord].beam || beamRow[j] {
				curCell.beam = true
			}

			up := Coord{X: coord.X, Y: coord.Y - 1}
			left := Coord{X: coord.X - 1, Y: coord.Y}
			right := Coord{X: coord.X + 1, Y: coord.Y}
			if r == '^' {
				curCell.splitter = true
				curCell.beam = false
				delete(beamRow, j)
				grid[coord] = curCell
				if grid[up].beam {
					answer++
					if !grid[left].splitter {
						l := grid[left]
						l.beam = true
						grid[left] = l
					}
					grid[right] = cell{coord: Coord{X: coord.X + 1, Y: coord.Y}, beam: true}
				}
				continue
			}
			if grid[up].beam {
				curCell.beam = true
				grid[coord] = curCell
			}
		}
	}
	/*
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[0]); j++ {
				coord := Coord{X: j, Y: i}
				if grid[coord].beam {
					fmt.Print("|")
				} else if grid[coord].splitter {
					fmt.Print("^")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	*/
	return answer
}

func part2(lines []string) int {
	answer := 0
	grid := map[Coord]*cell{}
	start := Coord{}

	for i, line := range lines {
		for j, r := range line {
			coord := Coord{X: j, Y: i}
			curCell := &cell{coord: coord}
			if r == 'S' {
				curCell.beam = true
				grid[coord] = curCell
				start = coord
				continue
			}

			if r == '^' {
				curCell.splitter = true
				grid[coord] = curCell
				continue
			}

			grid[coord] = curCell
		}
	}

	answer = travel(grid, start, len(lines))

	return answer
}

func travel(grid map[Coord]*cell, coord Coord, max int) int {
	if coord.Y >= max {
		return 1
	}
	if grid[coord] == nil {
		return 0
	}
	if grid[coord].downs > 0 {
		return grid[coord].downs
	}
	world := 0
	if grid[coord].splitter {
		world += travel(grid, Coord{X: coord.X - 1, Y: coord.Y}, max)
		world += travel(grid, Coord{X: coord.X + 1, Y: coord.Y}, max)
	} else {
		world += travel(grid, Coord{X: coord.X, Y: coord.Y + 1}, max)
	}
	c := grid[coord]
	c.downs = world
	grid[coord] = c
	return world
}
