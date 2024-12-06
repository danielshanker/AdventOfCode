package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 161
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day3/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 48
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day3/input.txt")
		fmt.Printf("day3 Answer 1 : %d\n", part1(input))
		fmt.Printf("day3 Answer 2 : %d\n", part2(input))
	}

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
