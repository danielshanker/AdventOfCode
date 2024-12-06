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
		expectedAnswer := 2
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day2/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 4
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day2/input.txt")
		fmt.Printf("day2 Answer 1 : %d\n", part1(input))
		fmt.Printf("day2 Answer 2 : %d\n", part2(input))
	}

}

const inc = 1
const dec = -1

func part1(lines []string) int {
	answer := 0

	for _, line := range lines {
		levels := strings.Fields(line)

		if checkLevels(levels, -1) {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	for _, line := range lines {
		levels := strings.Fields(line)

		safe := checkLevels(levels, -1)
		if !safe {
			for i := range levels {
				if checkLevels(levels, i) {
					answer++
					break
				}
			}
		} else {
			answer++
		}
	}

	return answer
}

func checkLevels(levels []string, ignoreIndex int) bool {
	lastLevel := S2i(levels[0])
	if ignoreIndex == 0 {
		lastLevel = S2i(levels[1])
	}
	dir := 0
	unsafe := false
	for i, level := range levels {
		if i == 0 || i == ignoreIndex || (ignoreIndex == 0 && i == 1) {
			continue
		}
		l := S2i(level)
		if dir == 0 {
			if l > lastLevel {
				dir = inc
			} else if l < lastLevel {
				dir = dec
			} else {
				unsafe = true
				break
			}
		}
		if dir == inc {
			if l <= lastLevel {
				unsafe = true
				break
			}
			if l-lastLevel > 3 {
				unsafe = true
				break
			}
			lastLevel = l
		}
		if dir == dec {
			if l >= lastLevel {
				unsafe = true
				break
			}
			if lastLevel-l > 3 {
				unsafe = true
				break
			}
			lastLevel = l
		}
	}

	return !unsafe
}
