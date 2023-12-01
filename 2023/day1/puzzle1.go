package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day1/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day1/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0

	for _, line := range lines {
		first := 0
		last := 0
		for _, cha := range line {
			num, err := strconv.Atoi(string(cha))
			if err == nil {
				first = num
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			cha := line[i]
			num, err := strconv.Atoi(string(cha))
			if err == nil {
				last = num
				break
			}
		}
		value := first*10 + last
		answer += value
	}

	fmt.Printf("day1 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0
	for _, line := range lines {
		first := 0
		last := 0
		newLine := ""
		for _, cha := range line {
			num, err := strconv.Atoi(string(cha))
			if err == nil {
				first = num
				break
			}
			newLine += string(cha)
			number, found := convertStringToNum(newLine)
			if found {
				first = number
				break
			}
		}
		newLine = ""
		for i := len(line) - 1; i >= 0; i-- {
			cha := line[i]
			num, err := strconv.Atoi(string(cha))
			if err == nil {
				last = num
				break
			}
			newLine = string(cha) + newLine
			number, found := convertStringToNum(newLine)
			if found {
				last = number
				break
			}
		}
		value := first*10 + last
		answer += value
	}

	fmt.Printf("day1 Answer 2 : %d\n", answer)
}

func convertStringToNum(line string) (int, bool) {
	conv := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for snum, inum := range conv {
		if strings.Contains(line, snum) {
			return inum, true
		}
	}
	return 0, false
}
