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
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day17/sample.txt")
		part1(sample[0])
		part2(sample[0])
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day17/input.txt")
		part1(input[0])
		part2(input[0])
	}

}

var shapes = []string{
	"-",
	"+",
	"j",
	"l",
	"o",
}

type rock struct {
	x map[int]int
	y map[int]int
}

func part1(line string) {
	answer := 0
	shapeIndex := 0
	jetIndex := 0
	var cave [7]map[int]bool
	for i := 0; i < 7; i++ {
		cave[i] = map[int]bool{}
	}

	for i := 0; i < 2022; i++ {
		curShape := shapes[shapeIndex%5]

		shape := getShape(curShape, answer)

		for {
			move := string(line[jetIndex%len(line)])
			jetIndex++
			blocked := false
			if move == "<" {
				for j := 0; j < len(shape.x); j++ {
					if shape.x[j]-1 < 0 || cave[shape.x[j]-1][shape.y[j]] {
						blocked = true
						break
					}
				}
				if !blocked {
					for j := 0; j < len(shape.x); j++ {
						shape.x[j]--
					}

				}
			} else {
				for j := 0; j < len(shape.x); j++ {
					if shape.x[j]+1 >= 7 || cave[shape.x[j]+1][shape.y[j]] {
						blocked = true
						break
					}
				}
				if !blocked {
					for j := 0; j < len(shape.x); j++ {
						shape.x[j]++
					}
				}
			}
			hitBottom := false
			for j := 0; j < len(shape.y); j++ {
				if shape.y[j]-1 <= 0 || cave[shape.x[j]][shape.y[j]-1] {
					hitBottom = true
					break
				}
			}
			if hitBottom {
				break
			}
			for j := 0; j < len(shape.y); j++ {
				shape.y[j]--
			}

		}
		for j := 0; j < len(shape.y); j++ {
			if shape.y[j] > answer {
				answer = shape.y[j]
			}

			cave[shape.x[j]][shape.y[j]] = true
		}
		shapeIndex++
	}

	fmt.Printf("Answer 1 : %d\n", answer)
}

func part2(line string) {
	answer := 0
	shapeIndex := 0
	jetIndex := 0
	heights := map[int][]int{}
	heights2 := map[int]int{}
	var pStart, pLength int
	var cave [7]map[int]bool
	for i := 0; i < 7; i++ {
		cave[i] = map[int]bool{}
	}
	for i := 0; i < 1000000000000; i++ {
		curShape := shapes[shapeIndex%5]

		shape := getShape(curShape, answer)

		for {
			move := string(line[jetIndex%len(line)])
			jetIndex++
			blocked := false
			if move == "<" {
				for j := 0; j < len(shape.x); j++ {
					if shape.x[j]-1 < 0 || cave[shape.x[j]-1][shape.y[j]] {
						blocked = true
						break
					}
				}
				if !blocked {
					for j := 0; j < len(shape.x); j++ {
						shape.x[j]--
					}

				}
			} else {
				for j := 0; j < len(shape.x); j++ {
					if shape.x[j]+1 >= 7 || cave[shape.x[j]+1][shape.y[j]] {
						blocked = true
						break
					}
				}
				if !blocked {
					for j := 0; j < len(shape.x); j++ {
						shape.x[j]++
					}
				}
			}
			hitBottom := false
			for j := 0; j < len(shape.y); j++ {
				if shape.y[j]-1 <= 0 || cave[shape.x[j]][shape.y[j]-1] {
					hitBottom = true
					break
				}
			}
			if hitBottom {
				break
			}
			for j := 0; j < len(shape.y); j++ {
				shape.y[j]--
			}

		}
		for j := 0; j < len(shape.y); j++ {
			if shape.y[j] >= answer {
				answer = shape.y[j]
				heights[answer] = append(heights[answer], i)
				heights2[i] = answer
			}

			cave[shape.x[j]][shape.y[j]] = true
		}
		shapeIndex++
		if i > 10000 {
			pStart, pLength = lookForPattern(cave, i)
			if pStart > 0 && pLength > 0 {
				break
			}
		}
	}
	startCycle := heights[pStart]
	endCycle := heights[pStart+pLength-1]
	a := 1000000000000 - startCycle[0]
	cycles := a / (endCycle[0] - startCycle[0])
	extra := a%(endCycle[0]-startCycle[0]) - 1
	e := heights2[startCycle[0]+extra] - pStart
	answer = pStart + pLength*cycles + e

	fmt.Printf("Answer 2 : %d\n", answer)
}

func getShape(curShape string, height int) rock {

	var shape rock
	shape.x = map[int]int{}
	shape.y = map[int]int{}

	if curShape == "-" {
		/*
			0123
		*/
		shape.x[0] = 2
		shape.x[1] = 3
		shape.x[2] = 4
		shape.x[3] = 5

		shape.y[0] = height + 4
		shape.y[1] = height + 4
		shape.y[2] = height + 4
		shape.y[3] = height + 4
	} else if curShape == "+" {
		/*
				4
			   012
			    3
		*/
		shape.x[0] = 2
		shape.x[1] = 3
		shape.x[2] = 4
		shape.x[3] = 3
		shape.x[4] = 3

		shape.y[0] = height + 5
		shape.y[1] = height + 5
		shape.y[2] = height + 5
		shape.y[3] = height + 4
		shape.y[4] = height + 6
	} else if curShape == "j" {
		/*
		          4
		   		  3
		   		012
		*/

		shape.x[0] = 2
		shape.x[1] = 3
		shape.x[2] = 4
		shape.x[3] = 4
		shape.x[4] = 4

		shape.y[0] = height + 4
		shape.y[1] = height + 4
		shape.y[2] = height + 4
		shape.y[3] = height + 5
		shape.y[4] = height + 6
	} else if curShape == "l" {
		/*
		   3
		   2
		   1
		   0
		*/

		shape.x[0] = 2
		shape.x[1] = 2
		shape.x[2] = 2
		shape.x[3] = 2

		shape.y[0] = height + 4
		shape.y[1] = height + 5
		shape.y[2] = height + 6
		shape.y[3] = height + 7

	} else {
		/*
			23
			01
		*/

		shape.x[0] = 2
		shape.x[1] = 3
		shape.x[2] = 2
		shape.x[3] = 3

		shape.y[0] = height + 4
		shape.y[1] = height + 4
		shape.y[2] = height + 5
		shape.y[3] = height + 5
	}

	return shape
}

func lookForPattern(cave [7]map[int]bool, index int) (int, int) {
	caveLines := []string{}
	s := ""
	for i := 1; i < index; i++ {
		for j := 0; j < 7; j++ {
			if cave[j][i] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += " "
		if i == 50 {
			caveLines = append(caveLines, s)
		}
		if i > 50 {
			s = s[8:]
			caveLines = append(caveLines, s)
		}
	}

	found := false
	patternStart := 0
	patternLength := 0
	for j := 0; j < len(caveLines); j++ {
		for i := 1; i < len(caveLines); i++ {
			a := caveLines[i]
			b := caveLines[j]
			if a == b && i != j {
				if found {
					patternLength = i - patternStart
					return patternStart, patternLength
				}
				found = true
				patternStart = i
				continue
			}
		}
	}
	return 0, 0
}
