package main

import (
	"flag"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, dayNum, part1, part2, 0, 0)

}

func part1(lines []string) int {
	answer := 0

	return answer
}

func part2(lines []string) int {
	answer := 0

	return answer
}
