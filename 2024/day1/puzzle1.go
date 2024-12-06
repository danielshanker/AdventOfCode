package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"
	"utils"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 0
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day1/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", 11, answer1))
		}
		expectedAnswer = 31
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2024/day1/input.txt")
		fmt.Printf("day1 Answer 1 : %d\n", part1(input))
		fmt.Printf("day1 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0
	left := []int{}
	right := []int{}

	for _, line := range lines {
		vals := strings.Fields(line)
		left = append(left, utils.S2i(vals[0]))
		right = append(right, utils.S2i(vals[1]))
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		answer += int(math.Abs(float64(left[i] - right[i])))
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	left := []int{}
	right := []int{}
	for _, line := range lines {
		vals := strings.Fields(line)
		left = append(left, utils.S2i(vals[0]))
		right = append(right, utils.S2i(vals[1]))
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		target := left[i]
		count := 0
		for j := 0; j < len(right); j++ {
			if right[j] > target {
				break
			}
			if right[j] == target {
				count++
			}
		}
		answer += target * count
	}

	return answer
}
