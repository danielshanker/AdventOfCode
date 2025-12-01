package main

import (
	"flag"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 1, part1, part2, 3, 6)

}

func part1(lines []string) int {
	answer := 0

	curVal := 50

	for _, line := range lines {
		dir := line[:1]
		val := S2i(line[1:])
		if dir == "R" {
			curVal += val
			curVal = curVal % 100
		} else {
			curVal -= val
			curVal = curVal % 100
			if curVal < 0 {
				curVal = 100 + curVal
			}
		}
		if curVal == 0 {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	curVal := 50

	for _, line := range lines {
		dir := line[:1]
		val := S2i(line[1:])
		if val > 100 {
			answer += val / 100
			val = val % 100
		}
		if dir == "R" {
			curVal += val
			if curVal >= 100 {
				answer += curVal / 100
			}
			curVal = curVal % 100
		} else {
			isZero := curVal == 0
			curVal -= val
			if curVal <= 0 && !isZero {
				answer += curVal/-100 + 1
			}
			curVal = curVal % 100
			if curVal < 0 {
				curVal = 100 + curVal
			}
		}
	}

	return answer
}
