package main

import (
	"flag"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 12, part1, part2, 0, 0)

}

func part1(lines []string) int {
	answer := 0

	gifts := []int{}
	gift := 0
	for _, line := range lines {
		if len(gifts) <= 5 {
			if len(line) == 0 {
				gifts = append(gifts, gift)
				gift = 0
				continue
			}
			for _, r := range line {
				if r == '#' {
					gift++
					continue
				}
			}
			continue
		}
		if len(line) == 0 {
			continue
		}

		s := strings.Split(line, ": ")
		m := strings.Split(s[0], "x")
		area := S2i(m[0]) * S2i(m[1])
		gs := strings.Fields(s[1])
		total := 0
		for i, g := range gs {
			total += S2i(g) * gifts[i]
		}

		// if the area-total is negative, there is definitely not enough space.
		// After checking that, I looked at the output of the differences and they were all huge, indidcating that there was more than enough space for the presents
		if area-total >= 0 {
			answer++
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	return answer
}
