package main

import (
	"flag"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 2, part1, part2, 2, 4)
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
