package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 0
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/dayx/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 0
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/dayx/input.txt")
		fmt.Printf("dayx Answer 1 : %d\n", part1(input))
		fmt.Printf("dayx Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0

	return answer
}

func part2(lines []string) int {
	answer := 0

	return answer
}
