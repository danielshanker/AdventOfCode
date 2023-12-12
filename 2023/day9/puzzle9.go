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
		expectedAnswer := 114
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day9/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 2
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day9/input.txt")
		fmt.Printf("day9 Answer 1 : %d\n", part1(input))
		fmt.Printf("day9 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0

	patterns := [][]int{}

	for _, line := range lines {
		stringPattern := strings.Fields(line)
		pattern := []int{}
		for _, num := range stringPattern {
			pattern = append(pattern, S2i(num))
		}
		patterns = append(patterns, pattern)
	}

	for _, pattern := range patterns {
		sequences := [][]int{}
		newPattern := pattern
		sequences = append(sequences, pattern)
		for {
			nextPattern := []int{}
			allZeros := true
			for _, val := range newPattern {
				if val != 0 {
					allZeros = false
					break
				}
			}
			if allZeros {
				break
			}
			for i, num := range newPattern {
				if i == len(newPattern)-1 {
					break
				}

				val := newPattern[i+1] - num
				nextPattern = append(nextPattern, val)
			}
			newPattern = nextPattern
			sequences = append(sequences, nextPattern)
		}

		for i := len(sequences) - 1; i >= 0; i-- {
			if i == 0 {
				answer += sequences[0][len(sequences[0])-1]
				break
			}
			lastVal := sequences[i][len(sequences[i])-1]
			sequences[i-1] = append(sequences[i-1], sequences[i-1][len(sequences[i-1])-1]+lastVal)
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	patterns := [][]int{}

	for _, line := range lines {
		stringPattern := strings.Fields(line)
		pattern := []int{}
		for _, num := range stringPattern {
			pattern = append(pattern, S2i(num))
		}
		patterns = append(patterns, pattern)
	}

	for _, pattern := range patterns {
		sequences := [][]int{}
		newPattern := pattern
		sequences = append(sequences, pattern)
		for {
			nextPattern := []int{}
			allZeros := true
			for _, val := range newPattern {
				if val != 0 {
					allZeros = false
					break
				}
			}
			if allZeros {
				break
			}
			for i, num := range newPattern {
				if i == len(newPattern)-1 {
					break
				}

				val := newPattern[i+1] - num
				nextPattern = append(nextPattern, val)
			}
			newPattern = nextPattern
			sequences = append(sequences, nextPattern)
		}

		for i := len(sequences) - 1; i >= 0; i-- {
			if i == 0 {
				answer += sequences[0][0]
				break
			}
			firstValue := sequences[i][0]
			newVal := sequences[i-1][0] - firstValue
			sequences[i-1] = append([]int{newVal}, sequences[i-1]...)
		}
	}

	return answer
}
