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
		expectedAnswer := 8
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day10/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		sample = ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day10/sample2.txt")
		expectedAnswer = 10
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day10/input.txt")
		fmt.Printf("day10 Answer 1 : %d\n", part1(input))
		fmt.Printf("day10 Answer 2 : %d\n", part2(input))
	}

}

type pipePart struct {
	value  string
	inloop bool
	dir    string
	count  int
}

func part1(lines []string) int {
	answer := 0
	startX, startY, lastStep, coords := fillGrid(lines)
	for y, line := range lines {
		xCoord := map[int]pipePart{}
		for x, char := range line {
			pipeVal := string(char)
			if pipeVal == "S" {
				startX = x
				startY = y
				l := false
				r := false
				u := false
				d := false
				if x < len(line)-1 {
					if string(line[x+1]) == "-" || string(line[x+1]) == "7" || string(line[x+1]) == "J" {
						r = true
					}
				}
				if x > 0 {
					if string(line[x-1]) == "-" || string(line[x-1]) == "F" || string(line[x-1]) == "L" {
						l = true
					}
				}
				if y < len(lines)-1 {
					if string(lines[y+1][x]) == "|" || string(lines[y+1][x]) == "L" || string(lines[y+1][x]) == "J" {
						d = true
					}
				}
				if y > 0 {
					if string(lines[y-1][x]) == "|" || string(lines[y-1][x]) == "F" || string(lines[y-1][x]) == "7" {
						u = true
					}
				}
				if u && d {
					pipeVal = "|"
					lastStep = "u"
				}
				if d && r {
					pipeVal = "F"
					lastStep = "u"
				}
				if d && l {
					pipeVal = "7"
					lastStep = "u"
				}
				if u && r {
					pipeVal = "L"
					lastStep = "d"
				}
				if u && l {
					pipeVal = "J"
					lastStep = "d"
				}
				if l && r {
					pipeVal = "-"
					lastStep = "r"
				}

			}
			pipe := pipePart{
				value: pipeVal,
			}
			xCoord[x] = pipe
		}
		coords[y] = xCoord
	}

	curX := startX
	curY := startY
	steps := 1
	for {
		curPipe := coords[curY][curX]
		curPipeVal := curPipe.value
		curPipe.inloop = true
		curPipe.count = steps
		steps++

		if curPipeVal == "|" && lastStep == "d" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "|" && lastStep == "u" {
			curY--
			lastStep = "u"
			curPipe.dir = "u"
		} else if curPipeVal == "-" && lastStep == "r" {
			curX++
			lastStep = "r"
			curPipe.dir = "r"
		} else if curPipeVal == "-" && lastStep == "l" {
			curX--
			lastStep = "l"
			curPipe.dir = "l"
		} else if curPipeVal == "7" && lastStep == "r" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "7" && lastStep == "u" {
			curX--
			lastStep = "l"
			curPipe.dir = "l"
		} else if curPipeVal == "F" && lastStep == "l" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "F" && lastStep == "u" {
			curX++
			lastStep = "r"
			curPipe.dir = "r"
		} else if curPipeVal == "J" && lastStep == "d" {
			curX--
			lastStep = "l"
			curPipe.dir = "l"
		} else if curPipeVal == "J" && lastStep == "r" {
			curY--
			lastStep = "u"
			curPipe.dir = "u"
		} else if curPipeVal == "L" && lastStep == "d" {
			curX++
			lastStep = "r"
			curPipe.dir = "r"
		} else if curPipeVal == "L" && lastStep == "l" {
			curY--
			lastStep = "u"
			curPipe.dir = "u"
		} else {
			fmt.Println("FAIL")
		}

		if curX == startX && curY == startY {
			break
		}

	}
	answer = steps / 2
	return answer
}

func part2(lines []string) int {
	answer := 0
	coords := map[int]map[int]pipePart{}
	startX := 0
	startY := 0
	lastStep := ""

	curX := startX
	curY := startY
	steps := 1
	for {
		lastX, lastY := curX, curY
		curPipe := coords[curY][curX]
		curPipeVal := curPipe.value
		curPipe.inloop = true
		curPipe.count = steps
		steps++

		if curPipeVal == "|" && lastStep == "d" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "|" && lastStep == "u" {
			curY--
			lastStep = "u"
			curPipe.dir = "u"
		} else if curPipeVal == "-" && lastStep == "r" {
			curX++
			lastStep = "r"
		} else if curPipeVal == "-" && lastStep == "l" {
			curX--
			lastStep = "l"
		} else if curPipeVal == "7" && lastStep == "r" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "7" && lastStep == "u" {
			curX--
			lastStep = "l"
			curPipe.dir = "u"
		} else if curPipeVal == "F" && lastStep == "l" {
			curY++
			lastStep = "d"
			curPipe.dir = "d"
		} else if curPipeVal == "F" && lastStep == "u" {
			curX++
			lastStep = "r"
			curPipe.dir = "u"
		} else if curPipeVal == "J" && lastStep == "d" {
			curX--
			lastStep = "l"
		} else if curPipeVal == "J" && lastStep == "r" {
			curY--
			lastStep = "u"
		} else if curPipeVal == "L" && lastStep == "d" {
			curX++
			lastStep = "r"
		} else if curPipeVal == "L" && lastStep == "l" {
			curY--
			lastStep = "u"
		} else {
			fmt.Println("FAIL")
		}
		coords[lastY][lastX] = curPipe
		if curX == startX && curY == startY {
			break
		}
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			count := 0
			if coords[y][x].inloop {
				continue
			}
			for i := x; i < len(lines[y]); i++ {
				curPipe := coords[y][i]
				if !curPipe.inloop {
					continue
				}
				if curPipe.dir == "u" {
					count++
				}
				if curPipe.dir == "d" {
					count--
				}
			}
			if count != 0 {
				answer++
			}
		}
	}

	return answer
}

func fillGrid(lines []string) (int, int, string, map[int]map[int]pipePart) {
	startX := 0
	startY := 0
	coords := map[int]map[int]pipePart{}
	lastStep := ""
	for y, line := range lines {
		xCoord := map[int]pipePart{}
		for x, char := range line {
			pipeVal := string(char)
			if pipeVal == "S" {
				startX = x
				startY = y
				l := false
				r := false
				u := false
				d := false
				if x < len(line)-1 {
					if string(line[x+1]) == "-" || string(line[x+1]) == "7" || string(line[x+1]) == "J" {
						r = true
					}
				}
				if x > 0 {
					if string(line[x-1]) == "-" || string(line[x-1]) == "F" || string(line[x-1]) == "L" {
						l = true
					}
				}
				if y < len(lines)-1 {
					if string(lines[y+1][x]) == "|" || string(lines[y+1][x]) == "L" || string(lines[y+1][x]) == "J" {
						d = true
					}
				}
				if y > 0 {
					if string(lines[y-1][x]) == "|" || string(lines[y-1][x]) == "F" || string(lines[y-1][x]) == "7" {
						u = true
					}
				}
				if u && d {
					pipeVal = "|"
					lastStep = "u"
				}
				if d && r {
					pipeVal = "F"
					lastStep = "u"
				}
				if d && l {
					pipeVal = "7"
					lastStep = "u"
				}
				if u && r {
					pipeVal = "L"
					lastStep = "d"
				}
				if u && l {
					pipeVal = "J"
					lastStep = "d"
				}
				if l && r {
					pipeVal = "-"
					lastStep = "r"
				}

			}
			pipe := pipePart{
				value: pipeVal,
			}
			xCoord[x] = pipe
		}
		coords[y] = xCoord
	}

	return startX, startY, lastStep, coords
}
