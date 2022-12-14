package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day14/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day14/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0
	cave, lowPoint := buildCave(lines, 1)

	answer = fall(cave, lowPoint, 1)

	fmt.Printf("Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0

	cave, lowPoint := buildCave(lines, 2)

	answer = fall(cave, lowPoint, 2)

	fmt.Printf("Answer 2 : %d\n", answer)
}
func fall(cave map[int]map[int]int, lowPoint int, part int) int {
	sandCount := 0
	for {
		sandCount++
		x := 500
		y := 0
		for j := 0; j < lowPoint+1; j++ {
			if j == lowPoint || cave[500][0] != 0 {
				return sandCount - 1
			}
			down := y + 1
			left := x - 1
			right := x + 1

			if _, ok := cave[x]; !ok {
				cave[x] = map[int]int{}
			}
			if _, ok := cave[left]; !ok {
				cave[left] = map[int]int{}
			}
			if _, ok := cave[right]; !ok {
				cave[right] = map[int]int{}
			}

			if cave[x][down] == 0 {
				y++
				continue
			}
			if cave[left][down] == 0 {
				y++
				x--
				continue
			}
			if cave[right][down] == 0 {
				y++
				x++
				continue
			}
			cave[x][y] = 2
			break
		}
	}
}

func buildCave(lines []string, part int) (map[int]map[int]int, int) {
	cave := map[int]map[int]int{}
	lowPoint := 0
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			xy1 := strings.Split(coords[i], ",")
			xy2 := strings.Split(coords[i+1], ",")
			x1 := S2i(xy1[0])
			x2 := S2i(xy2[0])
			y1 := S2i(xy1[1])
			y2 := S2i(xy2[1])
			if x1 == x2 {
				min := int(math.Min(float64(y1), float64(y2)))
				max := int(math.Max(float64(y1), float64(y2)))
				if max > lowPoint {
					lowPoint = max
				}
				for j := min; j <= max; j++ {
					if _, ok := cave[x1]; !ok {
						cave[x1] = map[int]int{}
					}
					cave[x1][j] = 1
				}
			} else {
				if y1 > lowPoint {
					lowPoint = y1
				}
				min := int(math.Min(float64(x1), float64(x2)))
				max := int(math.Max(float64(x1), float64(x2)))
				for j := min; j <= max; j++ {
					if _, ok := cave[j]; !ok {
						cave[j] = map[int]int{}
					}
					cave[j][y1] = 1
				}

			}

		}
	}
	if part == 2 {

		for j := -1000; j <= 1000; j++ {
			if _, ok := cave[j]; !ok {
				cave[j] = map[int]int{}
			}
			cave[j][lowPoint+2] = 1
		}
		return cave, lowPoint + 2
	}

	return cave, lowPoint
}
