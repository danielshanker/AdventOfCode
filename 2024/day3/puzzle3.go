package main

import (
	"flag"
	"regexp"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 3, part1, part2, 161, 48)
}

func part1(lines []string) int {
	answer := 0

	validMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	line := strings.Join(lines, "\n")
	muls := validMul.FindAllString(line, -1)
	for _, mul := range muls {
		vals := strings.Split(mul, ",")
		a := S2i(vals[0][4:])
		b := S2i(vals[1][:len(vals[1])-1])
		answer += a * b
	}
	return answer
}

func part2(lines []string) int {
	answer := 0

	validMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	dont := "don't()"
	do := "do()"
	line := strings.Join(lines, "\n")
	splitByDont := strings.Split(line, dont)
	muls := validMul.FindAllString(splitByDont[0], -1)
	for _, mul := range muls {
		vals := strings.Split(mul, ",")
		a := S2i(vals[0][4:])
		b := S2i(vals[1][:len(vals[1])-1])
		answer += a * b
	}
	for i, donts := range splitByDont {
		if i == 0 {
			continue
		}
		splitByDo := strings.Split(donts, do)
		for j, dos := range splitByDo {
			if j == 0 {
				continue
			}
			muls := validMul.FindAllString(dos, -1)
			for _, mul := range muls {
				vals := strings.Split(mul, ",")
				a := S2i(vals[0][4:])
				b := S2i(vals[1][:len(vals[1])-1])
				answer += a * b
			}

		}

	}

	return answer
}
