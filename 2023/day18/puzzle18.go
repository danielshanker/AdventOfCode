package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 62
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day18/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer2 := int64(952408144115)
		answer2 := part2(sample)
		if expectedAnswer2 == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer2, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day18/input.txt")
		fmt.Printf("day18 Answer 1 : %d\n", part1(input))
		fmt.Printf("day18 Answer 2 : %d\n", part2(input))
	}

}

type tile struct {
	active bool
}

func part1(lines []string) int {
	answer := 0

	grid := map[int]map[int]*tile{}

	direction := []string{}
	moveCount := []int{}

	for _, line := range lines {
		splitLine := strings.Fields(line)
		direction = append(direction, splitLine[0])
		moveCount = append(moveCount, S2i(splitLine[1]))
	}

	curY := 0
	minY := 1000000000
	maxY := -1000000000
	curX := 0
	minX := 1000000000
	maxX := -1000000000

	for i := 0; i < len(direction); i++ {
		for j := 0; j < moveCount[i]; j++ {
			if direction[i] == "R" {
				curX++
			}
			if direction[i] == "L" {
				curX--
			}
			if direction[i] == "U" {
				curY--
			}
			if direction[i] == "D" {
				curY++
			}

			minX = int(math.Min(float64(curX), float64(minX)))
			minY = int(math.Min(float64(curY), float64(minY)))
			maxX = int(math.Max(float64(curX), float64(maxX)))
			maxY = int(math.Max(float64(curY), float64(maxY)))

			if _, ok := grid[curY]; !ok {
				grid[curY] = map[int]*tile{}
			}
			curTile := tile{
				active: true,
			}
			grid[curY][curX] = &curTile
		}

	}

	fill(grid, (maxX+minX)/2, minY+1)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if grid[i][j] == nil {
				grid[i][j] = &tile{}
			}
			if grid[i][j].active {
				answer++
			}
		}
	}

	return answer
}

type coord struct {
	x int64
	y int64
}

func part2(lines []string) int64 {
	answer := int64(0)

	direction := []string{}
	moveCount := []int{}

	for _, line := range lines {
		splitLine := strings.Fields(line)
		convLine := splitLine[2][2:]
		move := convLine[:len(convLine)-2]
		moveInt, _ := strconv.ParseInt(move, 16, 64)
		moveCount = append(moveCount, int(moveInt))

		d := convLine[5:]
		d = d[:1]

		if d == "0" {
			direction = append(direction, "R")
		} else if d == "1" {
			direction = append(direction, "D")
		} else if d == "2" {
			direction = append(direction, "L")
		} else if d == "3" {
			direction = append(direction, "U")
		}
	}

	curY := int64(0)
	curX := int64(0)

	coords := []coord{}
	perim := 0

	for i := 0; i < len(direction); i++ {
		perim += moveCount[i]
		if direction[i] == "R" {
			curX += int64(moveCount[i])
		}
		if direction[i] == "L" {
			curX -= int64(moveCount[i])
		}
		if direction[i] == "U" {
			curY -= int64(moveCount[i])
		}
		if direction[i] == "D" {
			curY += int64(moveCount[i])
		}
		curCoord := coord{
			x: curX,
			y: curY,
		}
		coords = append(coords, curCoord)
	}

	sum1 := int64(0)
	sum2 := int64(0)

	for i := 0; i < len(coords)-1; i++ {
		sum1 += int64(coords[i].x * coords[i+1].y)
		sum2 += int64(coords[i].y * coords[i+1].x)
	}

	answer = int64(math.Abs(float64(sum1+coords[0].x*coords[len(coords)-1].y-sum2-coords[0].y*coords[len(coords)-1].x)))/2 + int64(perim/2+1)

	return answer
}

func fill(grid map[int]map[int]*tile, curX, curY int) {
	if grid[curY][curX] == nil {
		grid[curY][curX] = &tile{}
	}
	if grid[curY][curX].active {
		return
	}
	grid[curY][curX] = &tile{active: true}
	fill(grid, curX-1, curY)
	fill(grid, curX+1, curY)
	fill(grid, curX, curY+1)
	fill(grid, curX, curY-1)
}
