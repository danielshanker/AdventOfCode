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
		expectedAnswer := 46
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day16/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 51
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day16/input.txt")
		fmt.Printf("day16 Answer 1 : %d\n", part1(input))
		fmt.Printf("day16 Answer 2 : %d\n", part2(input))
	}

}

type tile struct {
	value      string
	activated  bool
	activeDirs []string
}

type beamType struct {
	x   int
	y   int
	dir string
}

func part1(lines []string) int {
	answer := 0
	grid := makeGrid(lines)

	beams := map[int]*beamType{0: {
		x:   -1,
		y:   0,
		dir: "r",
	}}

	answer = moveBeams(beams, grid)

	return answer
}

func part2(lines []string) int {
	answer := 0

	grid := makeGrid(lines)

	for i := 0; i < len(lines); i++ {
		deactivateGrid(grid)
		beams := map[int]*beamType{0: {
			x:   -1,
			y:   i,
			dir: "r",
		}}

		newAnswer := moveBeams(beams, grid)
		if newAnswer > answer {
			answer = newAnswer
		}
	}
	for i := 0; i < len(lines); i++ {
		deactivateGrid(grid)
		beams := map[int]*beamType{0: {
			x:   len(lines[0]),
			y:   i,
			dir: "l",
		}}

		newAnswer := moveBeams(beams, grid)
		if newAnswer > answer {
			answer = newAnswer
		}
	}
	for i := 0; i < len(lines[0]); i++ {
		deactivateGrid(grid)
		beams := map[int]*beamType{0: {
			x:   i,
			y:   -1,
			dir: "d",
		}}

		newAnswer := moveBeams(beams, grid)
		if newAnswer > answer {
			answer = newAnswer
		}
	}
	for i := 0; i < len(lines[0]); i++ {
		deactivateGrid(grid)
		beams := map[int]*beamType{0: {
			x:   i,
			y:   len(lines),
			dir: "u",
		}}

		newAnswer := moveBeams(beams, grid)
		if newAnswer > answer {
			answer = newAnswer
		}
	}

	return answer
}

func makeGrid(lines []string) map[int]map[int]*tile {
	grid := map[int]map[int]*tile{}
	for y, line := range lines {
		xGrid := map[int]*tile{}
		for x, cur := range line {
			char := string(cur)
			newTile := tile{
				value: char,
			}
			xGrid[x] = &newTile
		}
		grid[y] = xGrid
	}
	return grid
}

func moveBeams(beams map[int]*beamType, grid map[int]map[int]*tile) int {
	newBeamID := 1

	for {
		if len(beams) <= 0 {
			break
		}
		for i, beam := range beams {
			if beam.dir == "r" {
				beam.x++
			}
			if beam.dir == "l" {
				beam.x--
			}
			if beam.dir == "d" {
				beam.y++
			}
			if beam.dir == "u" {
				beam.y--
			}
			if beam.y < 0 || beam.y >= len(grid) || beam.x < 0 || beam.x >= len(grid[0]) {
				delete(beams, i)
				continue
			}
			deleteBeam := false
			newTile := grid[beam.y][beam.x]
			for _, dirs := range newTile.activeDirs {
				if dirs == beam.dir {
					deleteBeam = true
					break
				}
			}
			if deleteBeam {
				delete(beams, i)
				continue
			}
			newTile.activated = true
			newTile.activeDirs = append(newTile.activeDirs, beam.dir)
			if newTile.value == "-" {
				if beam.dir == "u" || beam.dir == "d" {
					beam.dir = "r"
					newBeam := beamType{
						dir: "l",
						x:   beam.x,
						y:   beam.y,
					}
					beams[newBeamID] = &newBeam
					newBeamID++
				}
			}
			if newTile.value == "|" {
				if beam.dir == "r" || beam.dir == "l" {
					beam.dir = "u"
					newBeam := beamType{
						dir: "d",
						x:   beam.x,
						y:   beam.y,
					}
					beams[newBeamID] = &newBeam
					newBeamID++
				}
			}
			if newTile.value == "\\" {
				if beam.dir == "r" {
					beam.dir = "d"
				} else if beam.dir == "u" {
					beam.dir = "l"
				} else if beam.dir == "l" {
					beam.dir = "u"
				} else if beam.dir == "d" {
					beam.dir = "r"
				}
			}
			if newTile.value == "/" {
				if beam.dir == "r" {
					beam.dir = "u"
				} else if beam.dir == "u" {
					beam.dir = "r"
				} else if beam.dir == "l" {
					beam.dir = "d"
				} else if beam.dir == "d" {
					beam.dir = "l"
				}
			}
		}
	}
	returnVal := 0
	for _, yLine := range grid {
		for _, val := range yLine {
			if val.activated {
				returnVal++
			}
		}
	}
	return returnVal
}

func deactivateGrid(grid map[int]map[int]*tile) {
	for _, yLine := range grid {
		for _, t := range yLine {
			t.activated = false
			t.activeDirs = []string{}
		}
	}
}
