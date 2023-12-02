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
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day2/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day2/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	for _, line := range lines {
		a := strings.Split(line, ": ")
		b := strings.Split(a[0], " ")
		gameNum := S2i(b[1])
		answer += gameNum

		sets := strings.Split(a[1], "; ")
	out:
		for _, set := range sets {
			colours := strings.Split(set, ", ")
			for _, colour := range colours {
				counts := strings.Split(colour, " ")
				if counts[1] == "red" && S2i(counts[0]) > maxRed {
					answer -= gameNum
					break out
				}
				if counts[1] == "blue" && S2i(counts[0]) > maxBlue {
					answer -= gameNum
					break out
				}
				if counts[1] == "green" && S2i(counts[0]) > maxGreen {
					answer -= gameNum
					break out
				}
			}
		}
	}

	fmt.Printf("day2 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0
	for _, line := range lines {
		a := strings.Split(line, ": ")

		sets := strings.Split(a[1], "; ")
		colourMap := make(map[string]int)
		for _, set := range sets {
			colours := strings.Split(set, ", ")
			for _, colour := range colours {
				counts := strings.Split(colour, " ")
				if S2i(counts[0]) > colourMap[counts[1]] {
					colourMap[counts[1]] = S2i(counts[0])
				}
			}
		}
		power := colourMap["red"] * colourMap["blue"] * colourMap["green"]
		answer += power
	}

	fmt.Printf("day2 Answer 2 : %d\n", answer)
}
