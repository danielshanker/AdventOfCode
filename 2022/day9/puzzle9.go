package main

import (
	"flag"
	"fmt"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day9/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day9/input.txt")
		part1(input)
		part2(input)
	}

}

type pos struct {
	x int
	y int
}

func part1(lines []string) {
	answer1 := 1
	grid := map[int]map[int]bool{
		0: map[int]bool{
			0: true,
		},
	}
	hPos := pos{x: 0, y: 0}
	tPos := pos{x: 0, y: 0}

	for _, line := range lines {
		command := strings.Split(line, " ")
		dir := command[0]
		num := S2i(command[1])

		if dir == "R" {
			hPos.x += num
		}
		if dir == "L" {
			hPos.x -= num
		}
		if dir == "U" {
			hPos.y += num
		}
		if dir == "D" {
			hPos.y -= num
		}

		if hPos.x != tPos.x && hPos.y != tPos.y {
			if isTouching(hPos, tPos) {
				continue
			}
			xDist := hPos.x - tPos.x
			yDist := hPos.y - tPos.y

			if yDist == 1 {
				tPos.y++
			}
			if yDist == -1 {
				tPos.y--
			}
			if xDist == 1 {
				tPos.x++
			}
			if xDist == -1 {
				tPos.x--
			}

		}

		if hPos.x == tPos.x || hPos.y == tPos.y {
			if hPos.x == tPos.x && hPos.y == tPos.y {
				continue
			}
			if isTouching(hPos, tPos) {
				continue
			}

			xDist := hPos.x - tPos.x
			yDist := hPos.y - tPos.y

			if xDist > 1 {
				for i := 0; i < xDist-1; i++ {
					tPos.x++
					if _, ok := grid[tPos.x]; !ok {
						grid[tPos.x] = map[int]bool{}
					}
					if _, ok := grid[tPos.x][tPos.y]; !ok {
						answer1++
					}
					grid[tPos.x][tPos.y] = true
				}
				continue
			}
			if xDist < -1 {
				for i := 0; i < -1*xDist-1; i++ {
					tPos.x--
					if _, ok := grid[tPos.x]; !ok {
						grid[tPos.x] = map[int]bool{}
					}
					if _, ok := grid[tPos.x][tPos.y]; !ok {
						answer1++
					}
					grid[tPos.x][tPos.y] = true
				}
				continue
			}
			if yDist > 1 {
				for i := 0; i < yDist-1; i++ {
					tPos.y++
					if _, ok := grid[tPos.x]; !ok {
						grid[tPos.x] = map[int]bool{}
					}
					if _, ok := grid[tPos.x][tPos.y]; !ok {
						answer1++
					}
					grid[tPos.x][tPos.y] = true
				}
				continue
			}
			if yDist < -1 {
				for i := 0; i < -1*yDist-1; i++ {
					tPos.y--
					if _, ok := grid[tPos.x]; !ok {
						grid[tPos.x] = map[int]bool{}
					}
					if _, ok := grid[tPos.x][tPos.y]; !ok {
						answer1++
					}
					grid[tPos.x][tPos.y] = true
				}
				continue
			}
		}

	}

	fmt.Printf("Answer 1 : %d\n", answer1)
}

func printGrid(grid map[int]map[int]bool, maxX int, maxY int, knots []pos) {
	for j := maxX; j >= -5; j-- {
		for i := -21; i < maxY; i++ {
			if i == 0 && j == 0 {
				fmt.Print("s")
				continue
			}
			out := false
			for a, knot := range knots {
				if knot.x == i && knot.y == j {
					fmt.Print(a)
					out = true
				}
			}
			if out {
				continue
			}
			if _, ok := grid[i]; !ok {
				fmt.Print(".")
				continue
			}
			if _, ok := grid[i][j]; !ok {
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func isTouching(h pos, t pos) bool {

	for i := -1; i <= 1; i++ {
		x := h.x + i
		for j := -1; j <= 1; j++ {
			y := h.y + j
			if t.x == x && t.y == y {
				return true
			}
		}
	}

	return false
}

func part2(lines []string) {
	answer2 := 1

	grid := map[int]map[int]bool{
		0: map[int]bool{
			0: true,
		},
	}
	knots := []pos{
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
		pos{x: 0, y: 0},
	}

	//lastDir := ""
	for _, line := range lines {
		command := strings.Split(line, " ")
		dir := command[0]
		num := S2i(command[1])

		for k := 0; k < num; k++ {

			if dir == "R" {
				knots[0].x++
			}
			if dir == "L" {
				knots[0].x--
			}
			if dir == "U" {
				knots[0].y++
			}
			if dir == "D" {
				knots[0].y--
			}
			for a := 1; a < 10; a++ {
				hPos := &knots[a-1]
				tPos := &knots[a]

				if hPos.x != tPos.x && hPos.y != tPos.y {
					if isTouching(*hPos, *tPos) {
						continue
					}
					xDist := hPos.x - tPos.x
					yDist := hPos.y - tPos.y

					if xDist >= 1 {
						if yDist >= 1 {
							tPos.x++
							tPos.y++
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
							continue
						}
						if yDist <= -1 {
							tPos.x++
							tPos.y--
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
							continue
						}
					}
					if xDist <= -1 {
						if yDist >= 1 {
							tPos.x--
							tPos.y++
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
							continue
						}
						if yDist <= -1 {
							tPos.x--
							tPos.y--
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
							continue
						}
					}

				}

				if hPos.x == tPos.x || hPos.y == tPos.y {
					if hPos.x == tPos.x && hPos.y == tPos.y {
						continue
					}
					if isTouching(*hPos, *tPos) {
						continue
					}

					xDist := hPos.x - tPos.x
					yDist := hPos.y - tPos.y

					if xDist > 1 {
						for i := 0; i < xDist-1; i++ {
							tPos.x++
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
						}
						continue
					}
					if xDist < -1 {
						for i := 0; i < -1*xDist-1; i++ {
							tPos.x--
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
						}
						continue
					}
					if yDist > 1 {
						for i := 0; i < yDist-1; i++ {
							tPos.y++
							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
						}
						continue
					}
					if yDist < -1 {
						for i := 0; i < -1*yDist-1; i++ {
							tPos.y--

							if a == 9 {
								if _, ok := grid[tPos.x]; !ok {
									grid[tPos.x] = map[int]bool{}
								}
								if _, ok := grid[tPos.x][tPos.y]; !ok {
									answer2++
								}
								grid[tPos.x][tPos.y] = true
							}
						}
						continue
					}
				}
			}
		}
	}

	fmt.Printf("Answer 2 : %d\n", answer2)
}
