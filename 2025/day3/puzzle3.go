package main

import (
	"flag"
	"math"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 3, part1, part2, 357, 3121910778619)

}

func part1(lines []string) int {
	answer := 0

	for _, line := range lines {
		left := 0
		right := 0
		for i, r := range line {
			cur := S2i(string(r))
			if i < len(line)-1 {
				if cur > left {
					left = cur
					right = 0
					continue
				}
			}
			if cur > right {
				right = cur
			}
		}
		answer += left*10 + right
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	for _, line := range lines {
		vals := make([]int, 12)
		for i, r := range line {
			cur := S2i(string(r))
			for j := 0; j < 12; j++ {
				if checkVal(i, len(line), 11-j, cur, vals[j]) {
					vals[j] = cur
					for k := j + 1; k < 12; k++ {
						vals[k] = 0
					}
					break
				}
			}
		}
		add := 0
		for k := 0; k < 12; k++ {
			add += vals[k] * int(math.Pow10(11-k))
		}
		answer += add
	}

	return answer
}

func checkVal(i int, lineLength int, pos int, cur int, original int) bool {
	if i < lineLength-pos {
		if cur > original {
			return true
		}
	}
	return false
}
