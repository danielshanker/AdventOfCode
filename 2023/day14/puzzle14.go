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
		expectedAnswer := 136
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day14/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 64
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day14/input.txt")
		fmt.Printf("day14 Answer 1 : %d\n", part1(input))
		fmt.Printf("day14 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0

	grid := map[int]map[int]string{}
	for y, line := range lines {
		lineCoord := map[int]string{}
		for x, cur := range line {
			char := string(cur)
			lineCoord[x] = char
		}
		grid[y] = lineCoord
	}

	tilt(grid, "up")

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "O" {
				answer += (len(grid) - i)
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	grid := map[int]map[int]string{}
	for y, line := range lines {
		lineCoord := map[int]string{}
		for x, cur := range line {
			char := string(cur)
			lineCoord[x] = char
		}
		grid[y] = lineCoord
	}

	arrangements := []string{}
	weights := []int{}
	weights = append(weights, 0)
	arrangements = append(arrangements, "")
	patternStart := 0
	patternLength := 0
	for i := 0; i < 1000; i++ {
		tilt(grid, "up")
		tilt(grid, "left")
		tilt(grid, "down")
		tilt(grid, "right")
		arrangement := ""
		weight := 0
		for k := 0; k < len(grid); k++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[k][j] == "O" {
					weight += (len(grid) - k)
				}
				arrangement += grid[k][j]
			}
		}
		if patternStart == 0 {
			for j, arr := range arrangements {
				if arr == arrangement {
					patternStart = j
					patternLength = i - j + 1
					break
				}
			}
		}
		arrangements = append(arrangements, arrangement)
		weights = append(weights, weight)
		if patternStart != 0 && i == patternStart+patternLength*2+1 {
			break
		}
	}

	start := 1000000000 - patternStart - 1

	index := start%patternLength + patternStart + 1
	fmt.Println()
	fmt.Println(patternStart)
	fmt.Println(patternLength)
	fmt.Println(index)

	answer = weights[index]
	fmt.Println(weights)

	return answer
}

func tilt(grid map[int]map[int]string, dir string) map[int]map[int]string {
	if dir == "up" {
		for i := 0; i < len(grid); i++ {
			for y, slice := range grid {
				for x, char := range slice {
					if char == "O" {
						if grid[y-1][x] == "." {
							grid[y-1][x] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}
	}
	if dir == "down" {
		for i := 0; i < len(grid); i++ {
			for y := len(grid) - 1; y >= 0; y-- {
				slice := grid[y]
				for x, char := range slice {
					if char == "O" {
						if grid[y+1][x] == "." {
							grid[y+1][x] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}
	}
	if dir == "left" {
		for i := 0; i < len(grid); i++ {
			for x := len(grid[0]) - 1; x >= 0; x-- {
				for y := len(grid) - 1; y >= 0; y-- {
					char := grid[y][x]
					if char == "O" {
						if grid[y][x-1] == "." {
							grid[y][x-1] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}
	}
	if dir == "right" {
		for i := 0; i < len(grid); i++ {
			for x := 0; x < len(grid[0]); x++ {
				for y := len(grid) - 1; y >= 0; y-- {
					char := grid[y][x]
					if char == "O" {
						if grid[y][x+1] == "." {
							grid[y][x+1] = "O"
							grid[y][x] = "."
						}
					}
				}
			}
		}
	}
	return grid
}

func printGrid(grid map[int]map[int]string) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
