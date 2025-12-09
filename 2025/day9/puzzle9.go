package main

import (
	"flag"
	"sort"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 9, part1, part2, 50, 24)

}

func part1(lines []string) int {
	answer := 0
	coords := []Coord{}
	for _, line := range lines {
		s := strings.Split(line, ",")
		c := Coord{
			X: S2i(s[0]),
			Y: S2i(s[1]),
		}
		coords = append(coords, c)
	}

	max := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			area := Area(coords[i], coords[j])
			if area > max {
				max = area
			}
		}
	}
	answer = max
	return answer
}

type area struct {
	c1 Coord
	c2 Coord
	a  int
}

type edge struct {
	c1 Coord
	c2 Coord
}

func part2(lines []string) int {
	answer := 0
	coords := []Coord{}
	for _, line := range lines {
		s := strings.Split(line, ",")
		c := Coord{
			X: S2i(s[0]),
			Y: S2i(s[1]),
		}
		coords = append(coords, c)
	}
	areas := []area{}
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			a := Area(coords[i], coords[j])
			areas = append(areas, area{c1: coords[i], c2: coords[j], a: a})
		}
	}

	sort.Slice(areas, func(i, j int) bool {
		return areas[i].a > areas[j].a
	})

	edges := []Coord{}
	for i, c := range coords {
		nextCoord := coords[(i+1)%len(coords)]
		edges = append(edges, getEdge(c, nextCoord)...)
	}

	for _, a := range areas {
		found := false
		for _, e := range edges {
			if e.X < Max(a.c1.X, a.c2.X) && e.X > Min(a.c1.X, a.c2.X) && e.Y < Max(a.c1.Y, a.c2.Y) && e.Y > Min(a.c1.Y, a.c2.Y) {
				found = true
				break
			}
		}
		if !found {
			answer = a.a
			break
		}
	}

	return answer
}

func getEdge(c1, c2 Coord) []Coord {
	useX := c1.X == c2.X
	coords := []Coord{}
	if useX {
		for i := Min(c1.Y, c2.Y); i < Max(c1.Y, c2.Y)+1; i++ {
			coords = append(coords, Coord{X: c1.X, Y: i})
		}
	} else {
		for i := Min(c1.X, c2.X); i < Max(c1.X, c2.X)+1; i++ {
			coords = append(coords, Coord{X: i, Y: c1.Y})
		}
	}
	return coords
}

// flood fill didn't work, too slow, keeping it here as a mark of my hubris
/*
func floodFill(grid [][]*cell, x, y int) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return
	}

	c := grid[y][x]

	if c.visited || c.full {
		return
	}

	c.visited = true
	c.full = true

	floodFill(grid, x+1, y)
	floodFill(grid, x-1, y)
	floodFill(grid, x, y+1)
	floodFill(grid, x, y-1)
}

func checkSquare(grid [][]*cell, c1, c2 Coord) bool {
	for y := Min(c1.Y, c2.Y); y < Max(c1.Y, c2.Y)+1; y++ {
		for x := Min(c1.Y, c2.Y); x < Max(c1.Y, c2.Y)+1; x++ {
			if grid[y][x].full {
				return false
			}
		}
	}
	return true
}
*/
