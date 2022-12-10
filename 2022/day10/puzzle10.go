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
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day10/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day10/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0

	cycle := 0
	x := 1

	for _, line := range lines {
		if line == "noop" {
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				answer1 += x * cycle
			}
			continue
		}

		inst := strings.Split(line, " ")
		val := S2i(inst[1])

		cycle++
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			answer1 += x * cycle
		}
		cycle++
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			answer1 += x * cycle
		}
		x += val
	}
	fmt.Printf("Answer 1 : %d\n", answer1)
}

func part2(lines []string) {
	answer2 := 0
	cycle := 0
	x := 1
	pos := 0

	for _, line := range lines {
		if line == "noop" {
			cycle++
			if x == pos || x == pos-1 || x == pos+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			pos++
			if pos == 40 {
				pos = 0
				fmt.Println()
			}
			continue
		}

		inst := strings.Split(line, " ")
		val := S2i(inst[1])

		cycle++
		if x == pos || x == pos-1 || x == pos+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		pos++
		if pos == 40 {
			pos = 0
			fmt.Println()
		}
		cycle++
		if x == pos || x == pos-1 || x == pos+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		pos++
		if pos == 40 {
			pos = 0
			fmt.Println()
		}
		x += val
	}
	fmt.Printf("Answer 2 : %d\n", answer2)
}
