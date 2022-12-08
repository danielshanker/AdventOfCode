package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day8/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day8/input.txt")
		part1(input)
		part2(input)
	}

}

type tree struct {
	size int
	seen bool
}

func part1(lines []string) {
	answer1 := 0

	var grid [][]tree

	for _, line := range lines {
		var l []tree
		for _, r := range line {
			char := string(r)
			t := tree{
				size: S2i(char),
				seen: false,
			}
			l = append(l, t)
		}
		grid = append(grid, l)
	}

	for i := 0; i < len(grid); i++ {
		tallest := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].size > tallest {
				tallest = grid[i][j].size
				if !grid[i][j].seen {
					answer1++
					grid[i][j].seen = true
				}
			}
		}
		tallest = -1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j].size > tallest {
				tallest = grid[i][j].size
				if !grid[i][j].seen {
					answer1++
					grid[i][j].seen = true
				}
			}
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		tallest := -1
		for j := 0; j < len(grid); j++ {
			if grid[j][i].size > tallest {
				tallest = grid[j][i].size
				if !grid[j][i].seen {
					answer1++
					grid[j][i].seen = true
				}
			}
		}
		tallest = -1
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[j][i].size > tallest {
				tallest = grid[j][i].size
				if !grid[j][i].seen {
					answer1++
					grid[j][i].seen = true
				}
			}
		}
	}

	fmt.Printf("Answer 1 : %d\n", answer1)
}

func part2(lines []string) {
	answer2 := 0
	var grid [][]tree

	for _, line := range lines {
		var l []tree
		for _, r := range line {
			char := string(r)
			t := tree{
				size: S2i(char),
				seen: false,
			}
			l = append(l, t)
		}
		grid = append(grid, l)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			scenic := look(grid, i, j, "u")
			scenic *= look(grid, i, j, "d")
			scenic *= look(grid, i, j, "l")
			scenic *= look(grid, i, j, "r")
			if scenic > answer2 {
				answer2 = scenic
			}
		}
	}

	fmt.Printf("Answer 2 : %d\n", answer2)
}

func look(grid [][]tree, x int, y int, dir string) int {
	size := grid[x][y].size
	trees := 0
	if dir == "u" {
		for i := x - 1; i >= 0; i-- {
			trees++
			if grid[i][y].size >= size {
				return trees
			}
		}
	}
	if dir == "d" {
		for i := x + 1; i < len(grid); i++ {
			trees++
			if grid[i][y].size >= size {
				return trees
			}
		}
	}
	if dir == "l" {
		for i := y - 1; i >= 0; i-- {
			trees++
			if grid[x][i].size >= size {
				return trees
			}
		}
	}
	if dir == "r" {
		for i := y + 1; i < len(grid[0]); i++ {
			trees++
			if grid[x][i].size >= size {
				return trees
			}
		}
	}
	return trees
}
